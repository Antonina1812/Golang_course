package http

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

type Post struct {
	ID       int
	Text     string
	Author   string
	Comments int
	Time     time.Time
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	s := ""
	for i := 0; i < 1000; i++ {
		p := &Post{ID: i, Text: "new post"}
		s += fmt.Sprintf("%#v", p)
	}
	w.Write([]byte(s))
}

func PprofPrint() {
	http.HandleFunc("/", handlePost)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

// в терминале подаём нагрузку: ab -t 20 -n 1000000 -c 5 http://127.0.0.1:8080/

// curl http://127.0.0.1:8080/debug/pprof/heap -o mem_out.txt
// curl http://127.0.0.1:8080/debug/pprof/profile?seconds=5 -o cpu_out.txt

// go tool pprof -svg -alloc_objects pprof_1.exe mem_out.txt > mem_ao.svg
// go tool pprof -svg pprof_1.exe cpu_out.txt > cpu.svg
