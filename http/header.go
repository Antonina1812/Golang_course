package http

import (
	"fmt"
	"net/http"
)

func handleHead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("RequestID", "d41a8cb77f90d434")

	fmt.Fprintln(w, "You browser is", r.UserAgent())
	fmt.Fprintln(w, "You accept", r.Header.Get("Accept"))
}

func HeaderPrint() {
	http.HandleFunc("/", handleHead)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
