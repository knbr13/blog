package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("secret-key")
	store = sessions.NewCookieStore(key)
)

func main() {
	m := map[interface{}]interface{}{
		"a": true,
		"b": false,
	}
	val, ok := m["a"].(bool)
	fmt.Println("val: ", val)
	fmt.Println("ok: ", ok)
	// http.HandleFunc("/secret/", secret)
	// http.HandleFunc("/login/", login)
	// http.HandleFunc("/logout/", logout)

	// http.ListenAndServe(":8080", nil)
}

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here

	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}
