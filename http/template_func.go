package http

import (
	"fmt"
	"html/template"
	"net/http"
)

func IsUserOdd(u *User) bool {
	return u.ID%2 != 0
}

func TemplateFuncPrint() {
	tmplFuncs := template.FuncMap{
		"OddUser": IsUserOdd,
	}

	tmpl, err := template.New("").Funcs(tmplFuncs).ParseFiles("func.html")
	if err != nil {
		panic(err)
	}

	users := []User{
		User{1, "Tonya", true},
		User{2, "Enot", false},
		User{3, "Go", true},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "func.html", struct {
			Users []User
		}{
			users, // экземпляр структуры
		})
	})

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
