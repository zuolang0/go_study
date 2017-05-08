package main
import "fmt"
func Factorial(x int)(result int){
    if x==0{
        result=1
    }else{
        result=x*Factorial(x-1)
    }
    return
}
func main() {
    var i int=15
    fmt.Printf("%d 的阶乘是%d \n",i,Factorial(i))
    //斐波那契数列
    for i=0;i<10;i++{
        fmt.Printf("%d\t",fibonaci(i))
    }
    fmt.Printf("\n")
    var sum int =17
    var count int =5
    var mean float32
    //类型转换
    //float32(sum)
    mean=float32(sum)/float32(count)
    fmt.Printf("mean 的值为:%f\n",mean)
}
/**
 * 斐波那契数列
 */
func fibonaci(n int)int{
    if n<2{
        return n
    }
    return fibonaci(n-2)+fibonaci(n-1)
}