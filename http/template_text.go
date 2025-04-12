package http

import (
	"fmt"
	"net/http"
	"text/template"
)

type tmplParams struct {
	URL     string
	Browser string
}

const EXAMPLE = `
Browser {{.Browser}}

you at {{.URL}}`

func handle3(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New(`example`)
	tmpl, _ = tmpl.Parse(EXAMPLE)

	params := tmplParams{
		URL:     r.URL.String(),
		Browser: r.UserAgent(),
	}

	tmpl.Execute(w, params) // выполнить
}

func TemplateTextPrint() {
	http.HandleFunc("/", handle3)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
