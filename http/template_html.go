package http

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	ID     int
	Name   string
	Active bool
}

func TemplateHTMLPrint() {
	tmpl := template.Must(template.ParseFiles("users.html"))

	users := []User{
		User{1, "Tonya", true},
		User{2, "<i>Enot<i>", false},
		User{3, "Go", true},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct {
			Users []User
		}{
			users, // экземпляр структуры
		})
	})

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
