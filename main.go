package main

import (
	"goweb/form"
	"goweb/model"
	"goweb/renderer"
	"goweb/repo"
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
		f := &form.NewPost{Post: &model.Post{}}
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

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.Handle("/", router)

	http.ListenAndServe(":3000", nil)
}
