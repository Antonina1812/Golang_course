package http

import (
	"fmt"
	"net/http"
)

func handler2(w http.ResponseWriter, r *http.Request) {
	myParam := r.URL.Query().Get("param")
	if myParam != "" {
		fmt.Fprintln(w, "`myParam` is", myParam)
	}

	key := r.FormValue("key")
	if key != "" {
		fmt.Fprintln(w, "`key` is", key)
	}
}

func GetPrint() {
	http.HandleFunc("/", handler2)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

// чтобы увидеть результат нужно в браузере зайти на страницу: http://127.0.0.1:8080/?param=1&key=2
