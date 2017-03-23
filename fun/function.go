package main
import "fmt"
/*可以返回两个值*/
func swap(x,y string)(string ,string){
    return y,x
}
func main(){
    a,b:=swap("LZ","YX")
    fmt.Println(a,b)
}