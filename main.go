package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TodoPageData struct {
	PageTitle string `json:"pageTitle"`
	Todos     []Todo `json:"todos"`
}

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("static/index.html"))
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/form/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("static/forms.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		tmpl.Execute(w, struct {
			Success bool
			ContactDetails
		}{true, ContactDetails{Email: details.Email, Subject: details.Subject, Message: details.Message}})
	})
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))
	http.ListenAndServe(":8080", nil)
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}
