package main

import (
	"fmt"
	"net/http"
)

func boy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I love you, boy!!!")
}
func girl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I love you, girl!!!")
}
func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I love you, everyone !!!")
}
func main() {
	/*使用handlefunc 直接将方法作为参数传递*/
	/*http.HandleFunc("/boy/", boy)
	  http.HandleFunc("/g/", girl)
	  http.HandleFunc("/", foo)
	  http.ListenAndServe(":8080", nil)*/
	//
	// http.Handler() 的第二个参数是要实现了 Handler 接口的类型
	// 可以通过类型强转来重新使用该函数来实现
	//运行失败
	http.Handle("/boy/", http.HandlerFunc(boy))
	http.Handle("/girl/", http.HandlerFunc(girl))
	http.Handle("/", http.HandlerFunc(foo))
	http.ListenAndServe(":8080", nil)
}
