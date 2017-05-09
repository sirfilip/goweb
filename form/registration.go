package form

import (
	"goweb/model"
	"net/http"

	"github.com/asaskevich/govalidator"
)

type UserRegisterForm struct {
	*model.User
	Errors []string
}

func RegistrationForm() *UserRegisterForm {
	return &UserRegisterForm{User: &model.User{}, Errors: make([]string, 0)}
}

func (u *UserRegisterForm) Submit(r *http.Request) bool {
	u.Username = r.FormValue("username")
	u.Email = r.FormValue("email")
	u.Password = r.FormValue("password")

	return u.Validate()
}

func (u *UserRegisterForm) Validate() bool {
	u.Errors = make([]string, 0)

	if !govalidator.IsEmail(u.Email) {
		u.Errors = append(u.Errors, "Email is not valid")
	}

	if len(u.Username) == 0 {
		u.Errors = append(u.Errors, "Username is required")
	}

	if len(u.Username) > 20 {
		u.Errors = append(u.Errors, "Username must not be longer then 20 chars")
	}

	if len(u.Password) == 0 {
		u.Errors = append(u.Errors, "Password is required")
	}

	return len(u.Errors) == 0
}
