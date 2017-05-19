package main

import (
	"fmt"
	"net/http"
)

type foo struct{}

func (f foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/boy":
		fmt.Fprintln(w, "I love you!!")
	case "/girl":
		fmt.Fprintln(w, "hehe")
	default:
		fmt.Fprintln(w, "Men would stop talking and women would shed tears when they see this.")
	}
}
func main() {
	var f foo
	http.ListenAndServe(":8080", f)
}
