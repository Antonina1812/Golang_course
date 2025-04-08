package http

import (
	"fmt"
	"net/http"
)

func PagesPrint() {
	http.HandleFunc("/page",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Single page:", r.URL.String())
		})

	http.HandleFunc("/pages/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Multiple pages:", r.URL.String())
		})

	http.HandleFunc("/", handler)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

// чтобы увидеть результат нужно зайти в браузер на страницу: 127.0.0.1:8080
// 127.0.0.1:8080/page
// 127.0.0.1:8080/pages/
