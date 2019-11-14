package middleware

import "net/http"

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

func HandlerFunc(fn http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		fn = m(fn)
	}
	return fn
}
