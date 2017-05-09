package main

import (
	"goweb/form"
	"goweb/middleware"
	"goweb/renderer"
	"goweb/repo"
	"goweb/service"
	"goweb/session"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sqlx.DB
	t  *template.Template
)

func init() {
	db = sqlx.MustConnect("sqlite3", ":memory:")
	t = template.Must(template.ParseGlob("templates/**/*.html"))

	db.MustExec(`CREATE TABLE IF NOT EXISTS posts (
		id integer primary key,
		title varchar,
		content varchar,
		published_at datetime
	)`)

	db.MustExec(`CREATE TABLE IF NOT EXISTS users (
		id integer primary key,
		username varchar,
		email varchar, 
		password varchar
	)`)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		posts, err := repo.PostRepository(db).Posts()
		if nil != err {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			renderer.New(w).Render("posts/index", map[string]interface{}{
				"Posts": posts,
			})
		}
	})

	router.HandleFunc("/posts/new", func(w http.ResponseWriter, r *http.Request) {
		f := form.NewPostForm()
		if r.Method == "POST" && f.Submit(r.FormValue("title"), r.FormValue("content")) {
			_, err := repo.PostRepository(db).Create(f.Post)
			if nil != err {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/", 301)
			return
		}
		renderer.New(w).Render("posts/new", f)
	})

	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		f := form.RegistrationForm()

		if r.Method == "POST" && f.Submit(r) {
			_, err := repo.UserRepository(db).Create(f.User)
			if nil != err {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/", 301)
			return
		}

		renderer.New(w).Render("users/new", f)
	})

	router.HandleFunc("/login", middleware.ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		f := form.LoginForm()

		if r.Method == "POST" && f.Submit(r) {
			user, err := repo.UserRepository(db).FindByUsername(f.Username)
			if nil != err {
				return err
			}
			if service.PasswordHasher().Check(user.Password, f.Password) {
				s, err := session.Get(r)
				if nil != err {
					return err
				}
				s.Values["auth_id"] = user.ID
				err = s.Save(r, w)
				if nil != err {
					return err
				}
				http.Redirect(w, r, "/", 302)
				return nil
			}
		}
		renderer.New(w).Render("login/login", f)
		return nil
	}))

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.Handle("/", router)

	http.ListenAndServe(":3000", nil)
}
