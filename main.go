package main

import (
	"fmt"
	"net/http"

	NETWORK "github.com/ricardo-lisboasantos/sample-middleware-in-go/network"
	OS "github.com/ricardo-lisboasantos/sample-middleware-in-go/os"
)

func main() {
	// Serve Static files
	router := NETWORK.CreateRouter()
	fmt.Println("Listening on port 80... \nFor Hello World: http://localhost/ (GET) \nFor Contact Form: http://localhost/api/contactme (POST)")
	OS.OpenBrowser("http://localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}

}
