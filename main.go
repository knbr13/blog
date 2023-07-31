package main

import (
	"html/template"
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

func main() {
	templ := template.Must(template.ParseFiles("static/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		templ.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}
