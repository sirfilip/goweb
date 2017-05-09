package middleware

import "net/http"

type routeHandler func(w http.ResponseWriter, r *http.Request) error

func ErrorHandler(f routeHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if nil != err {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
