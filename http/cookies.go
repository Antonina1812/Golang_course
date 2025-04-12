package http

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func mainPage(w http.ResponseWriter, r *http.Request) {
// 	session, err := r.Cookie("session_id")
// 	loggedIn := (err != http.ErrNoCookie)

// 	if loggedIn {
// 		fmt.Fprintln(w, `<a href="/logout">logout</a>`)
// 		fmt.Fprintln(w, "Welcome, "+session.Value)
// 	} else {
// 		fmt.Fprintln(w, `<a href="/login">login</a>`)
// 		fmt.Fprintln(w, "You need to login")
// 	}
// }

// func loginPage(w http.ResponseWriter, r *http.Request) {
// 	expiration := time.Now().Add(10 * time.Hour)
// 	cookie := http.Cookie{
// 		Name:    "session_id",
// 		Value:   "tonya",
// 		Expires: expiration,
// 	}
// 	http.SetCookie(w, &cookie)
// 	http.Redirect(w, r, "/", http.StatusFound)
// }

// func CookiesPrint() {

// }
