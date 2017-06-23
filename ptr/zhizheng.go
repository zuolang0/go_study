package main

//指针学习
import "fmt"

func main() {
	var a int = 10
	var ptr *int
	ptr = &a
	*ptr = 11
	fmt.Printf("ptr 的地址 %x\n", ptr)
	fmt.Printf("ptr 的值%d\n", *ptr)
	fmt.Printf("a 的值%d\n", a)
}
