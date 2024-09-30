package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Id          int
	Title       string
	Description string
	Done        bool
}

func main() {

	data := map[string][]Todo{
		"todos": {
			{Id: 1, Title: "Todo 1", Description: "Description 1", Done: false},
			{Id: 2, Title: "Todo 2", Description: "Description 2", Done: true},
			{Id: 3, Title: "Todo 3", Description: "Description 3", Done: false},
		},
	}

	todosHandler := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))

		templ.Execute(w, data)
	}

	addTodoHandler := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		todo := Todo{Id: len(data["todos"]) + 1, Title: title, Done: false}
		data["todos"] = append(data["todos"], Todo{Id: len(data["todos"]) + 1, Title: title, Done: false})
		templ := template.Must(template.ParseFiles("index.html"))
		templ.ExecuteTemplate(w, "todo-list-element", todo)
	}

	http.HandleFunc("/", todosHandler)
	http.HandleFunc("/add-todo", addTodoHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
