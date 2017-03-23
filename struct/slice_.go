package main
import "fmt"
/*func main() {
    //var  slice1 []int=make([]int,3)
    var number []int=make([]int,3,5)
    printSlice(number)
    if (number==nil) {
        fmt.Printf("空")
    }
}*/
func printSlice(x []int){
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
/*func main() {
    number :=[] int {1,2,3,4,5,6,7,8,9}
    printSlice(number)
    fmt.Printf("number==",number)
    fmt.Printf("\n")
    fmt.Printf("number[1:4]==",number[1:4])
    fmt.Printf("\n")
    fmt.Printf("number[:4]==",number[:4])
    fmt.Printf("\n")
    fmt.Printf("number[4:]==",number[4:])
    fmt.Printf("\n")
    number1 :=make([]int,1,5)
    fmt.Printf("number1==",number1)
    fmt.Printf("\n")
    printSlice(number1)
    number2 :=number[:2]
    printSlice(number2)
    number3 :=number[2:5]
    printSlice(number3)
}*/
func main() {
    var number []int
    printSlice(number)
    number=append(number,0)
    printSlice(number)
    number=append(number,1)
    printSlice(number)
    number=append(number,2,3,4)
    printSlice(number)
    number1 :=make([]int,len(number),(cap(number))*2)
    printSlice(number1)
    copy(number1,number)
    printSlice(number1)
}