package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key         = []byte("super-secret-key")
	sessionName = "gowebsession"
	store       *sessions.CookieStore
)

func init() {
	store = sessions.NewCookieStore(key)
}

func Get(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, sessionName)
}
