package main
import "fmt"
/*func main() {
    var a,b,c int
    a=10
    b=20
    c=a+b
    fmt.Printf("结果:a=%d,b=%d and c=%d\n",a,b,c)
}*/
var a int = 20;

func main() {
    //var a int =10
    var b int =20
    var c int =0
    fmt.Printf("main()函数中 a=%d\n",a);
    c=sum(a,b)
    fmt.Printf("main()函数中 c=%d\n",c);
}
func sum(a,b int) int {
    fmt.Printf("sum() 函数中 a = %d\n",  a);
    fmt.Printf("sum() 函数中 b = %d\n",  b);

    return a + b;
}