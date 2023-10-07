// Created by: Ricardo Santos
package network

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

// Defines a key and a store variable.
// The key variable is a byte slice that contains a secret key used for encryption and decryption of session data.
// The store variable is a sessions.CookieStore instance that is used to store session data in cookies.
// The store is initialized with the key variable, which ensures that the session data is encrypted and secure.
var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

// get secret
func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

// login
// Create a new session and set authenticated to true
func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

// logout
// Revoke users authentication
func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}
