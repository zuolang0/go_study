package main
import "fmt"
type Phone interface{
    call()
}
type Nokia struct{

}
func (nokia Nokia) call(){
    fmt.Println("I am Nokia,I can call you")
}
type IPone struct{

}
func (ipone IPone) call(){
    fmt.Println("I am ipone,I can call you")
}
func main() {
    var phone Phone
    phone= new(Nokia)
    phone.call()
    phone=new(IPone)
    phone.call()
}
/**
 * go 语言允许函数的重载,根据参数不同来写同名函数
 */