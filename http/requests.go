package http

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) { // w - то, куда записываем наш результат; r - сам запрос
	fmt.Fprintln(w, "hello") // это строчка работает так же, как и строчка ниже
	w.Write([]byte("!!!"))
}

func RequestsPrint() {
	http.HandleFunc("/", handler)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

// чтобы увидеть результат нужно зайти в браузер на страницу: 127.0.0.1:8080
