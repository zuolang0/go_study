package main
/*
func学习
 */
import "fmt"
func main(){
    var a int =100
    var b int =200
    var ret int
    ret=max(a,b)
    fmt.Printf("最大值是:%d\n",ret)
}
func max(num1,num2 int) int{
    /*
    max  函数名
    num1 num2 参数名
    int 参数类型
    int 返回类型
     */
    var result int
    if (num1>num2){
        result=num1
    }else{
        result=num2
    }
    return result
}