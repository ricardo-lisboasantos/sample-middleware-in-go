package middleware

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// // Check Auth Middleware
// func CheckAuth() Middleware {

// 	// Create a new Middleware
// 	return func(f http.HandlerFunc) http.HandlerFunc {

// 		// Define the http.HandlerFunc
// 		return func(w http.ResponseWriter, r *http.Request) {
// 			// Do middleware things
// 			// if r.Header.Get("Authorization") != "Bearer 1234567890" {
// 			// 	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
// 			// 	return
// 			// }

// 			session, _ := store.Get(r, "cookie-name")

// 			// Check if user is authenticated
// 			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
// 				return
// 			}

// 			// Call the next middleware/handler in chain
// 			f(w, r)
// 		}
// 	}
// }

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// enable CORS
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
