package main
import "fmt"
//go语言的error是一个接口类型
//type error interface{
// Error() string
//}
type DivideError struct{
    dividee int
    divider int
}
func (de *DivideError) Error() string{
    strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
    return fmt.Sprintf(strFormat,de.dividee)
}
//定义整除函数
func Divide(varDividee int,varDivider int)(result int,errorMsg string){
    if varDivider==0{
        dData:=DivideError{
            dividee:varDividee,
            divider:varDivider,
        }
        errorMsg=dData.Error()
        return
    }else{
        return varDividee/varDivider,""
    }
}
func main() {
    if result,errorMsg:=Divide(100,10);errorMsg==""{
        fmt.Println("100/10=",result)
    }
    if _,errorMsg:=Divide(100,0);errorMsg!="" {
        fmt.Println("errorMsg is:",errorMsg)
    }
}
//go 语言的大小引号使用有要求,main 函数中=="" 改为==''会报错