package middleware

import (
	"net/http"
)

func HttpMethod(m string) Middleware {
	return func(fn http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				br := http.StatusBadRequest
				http.Error(w, http.StatusText(br), br)
				return
			}
			fn(w, r)
		}
	}
}
