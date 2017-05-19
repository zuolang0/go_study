package main

import (
	"fmt"
	"net/http"
)

type boy struct{}

func (b boy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I love you!!")
}

type girl struct{}

func (g gril) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hehe!!")
}

type foo struct{}

func (f foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Men would stop talking and women would shed tears when they see this.")
}
func main() {
	var b boy
	var g girl
	var f foo
	//使用mux来为每个路由（操作）绑定页面
	// mux := http.NewServeMux()
	// mux.Handle("/boy/", b)
	// mux.Handle("/g/", g)
	// mux.Handle("/", f)
	// http.ListenAndServe(":8080", mux)

	//直接使用http.handle
	http.Handle("/boy/", b)
	http.Handle("/g/", g)
	http.Handle("/", f)
	http.ListenAndServe(":8080", nil)
}
