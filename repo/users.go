package repo

import (
	"goweb/model"
	"goweb/service"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	conn *sqlx.DB
}

func UserRepository(conn *sqlx.DB) *UserRepo {
	return &UserRepo{conn: conn}
}

func (u *UserRepo) Create(user *model.User) (*model.User, error) {
	password, err := service.PasswordHasher().Hash(user.Password)
	if nil != err {
		return user, err
	}
	user.Password = password

	result, err := u.conn.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", user.Username, user.Email, user.Password)
	if nil != err {
		return user, err
	}
	lastInsertId, err := result.LastInsertId()
	if nil != err {
		return user, err
	}
	user.ID = lastInsertId
	return user, nil
}

func (u *UserRepo) FindByUsername(username string) (*model.User, error) {
	user := &model.User{}
	if err := u.conn.Get(user, "SELECT * FROM users WHERE username = ?", username); nil != err {
		return user, err
	}
	return user, nil
}
