package network

import (
	MIDDLEWARE "github.com/ricardo-lisboasantos/sample-middleware-in-go/middleware"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {

	// Create router
	router := mux.NewRouter()

	// Web UI Routes
	router.HandleFunc("/", MIDDLEWARE.Chain(Index, MIDDLEWARE.Method("GET"), MIDDLEWARE.Logging()))
	router.HandleFunc("/contactme", MIDDLEWARE.Chain(ContactForm, MIDDLEWARE.Method("GET"), MIDDLEWARE.Logging()))

	// Serve Pages and Static Files
	// router.HandleFunc("/", MIDDLEWARE.Chain(Hello, MIDDLEWARE.Method("GET"), MIDDLEWARE.Logging()))
	router.HandleFunc("/static/", MIDDLEWARE.Chain(ServeStatic, MIDDLEWARE.Method("GET"), MIDDLEWARE.Logging()))

	// API Routes
	router.HandleFunc("/api/contactme", MIDDLEWARE.Chain(ContactForm, MIDDLEWARE.Method("POST"), MIDDLEWARE.Logging()))

	return router
}
