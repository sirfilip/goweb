package form

import "net/http"

type Login struct {
	Username string
	Password string
	Errors   []string
}

func LoginForm() *Login {
	return &Login{}
}

func (l *Login) Submit(r *http.Request) bool {
	l.Username = r.FormValue("username")
	l.Password = r.FormValue("password")
	return l.Validate()
}

func (l *Login) Validate() bool {
	l.Errors = make([]string, 0)
	if len(l.Username) == 0 {
		l.Errors = append(l.Errors, "Username cant be blank")
	}
	if len(l.Username) > 100 {
		l.Errors = append(l.Errors, "Username too long")
	}
	if len(l.Password) == 0 {
		l.Errors = append(l.Errors, "Password is required")
	}

	return len(l.Errors) == 0
}
