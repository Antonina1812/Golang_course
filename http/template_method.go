package http

import (
	"fmt"
	"net/http"
	"text/template"
)

func (u *User) PrintActive() string {
	if !u.Active {
		return ""
	}
	return "method says user " + u.Name + " active"
}

func TemplateMethodPrint() {
	tmpl, err := template.New("").ParseFiles("method.html")
	if err != nil {
		panic(err)
	}

	users := []User{
		User{1, "Tonya", true},
		User{2, "Enot", false},
		User{3, "Golang", true},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "method.html", struct {
			Users []User
		}{
			users, // экземпляр структуры
		})
	})

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
