// A simple web application example.
// Taken from: https://golang.org/doc/articles/wiki/
// https://ianmcloughlin.github.io :: 2017-09-13

package main

import (
	"fmt"
	"net/http"
)

func userinputhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Query().Get("value")) //.Path[1:])
}

func main() {

	// Adapted from: http://www.alexedwards.net/blog/serving-static-sites-with-go
	fs := http.FileServer(http.Dir("webFolder"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":3000", nil)
}
