package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	defer conn.Close()
	respond(conn)
}
func respond(conn net.Conn) {
	// 消息体
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title>Go example</title></head><body><strong>Hello World</strong></body></html>`
	// HTTP 协议及请求码
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	// 内容长度
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	// 内容类型
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
