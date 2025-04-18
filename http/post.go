package http

import (
	"fmt"
	"net/http"
)

var loginFormTmpl = []byte(`
<html>
	<body>
	<form action="/" method="post">
		Login: <input type="text" name="login">
		Password: <input type="password" name="password">
		<input type="submit" value="Login">
	</form>
	<body>
</html>
`)

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write(loginFormTmpl)
		return
	}

	inputLogin := r.FormValue("login")
	fmt.Fprintln(w, "you enter: ", inputLogin)
}

func PostPrint() {
	http.HandleFunc("/", mainPage)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
