package network

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	MIDDLEWARE "github.com/ricardo-lisboasantos/sample-middleware-in-go/middleware"
	WEBUI "github.com/ricardo-lisboasantos/sample-middleware-in-go/webui"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func ContactForm(w http.ResponseWriter, r *http.Request) {
	MIDDLEWARE.EnableCors(&w)
	vars := mux.Vars(r)
	name := vars["name"]       // the name
	email := vars["email"]     // the email
	message := vars["message"] // the message
	fmt.Println("Name: " + name + " Email: " + email + " Message: " + message)
	fmt.Fprintf(w, "Thanks for contacting me. I will get back to you shortly")
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func Index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, WEBUI.Index())
	tmpl, err := template.New("index").Parse(WEBUI.Index())
	check_error(err)
	err = tmpl.Execute(w, nil)
	check_error(err)

}

// http.HandleFunc("/secret", secret)
// http.HandleFunc("/login", login)
// http.HandleFunc("/logout", logout)
// http.ListenAndServe(":8080", nil)

func check_error(err error) {

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
