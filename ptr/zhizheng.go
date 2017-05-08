package main
import "fmt"
func main() {
    var a int =10
    var ptr *int
    ptr=&a
    fmt.Printf("ptr 的地址 %x\n",ptr)
    fmt.Printf("ptr 的值%d\n",*ptr)
}