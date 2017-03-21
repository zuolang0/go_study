/*package main
import "fmt"
func main() {
	fmt.Printf("Hello wrold")
	fmt.Printf("zl")
}*/
package main
/*package print*/

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

func Print(s string) {
    cs := C.CString(s)
    defer C.free(unsafe.Pointer(cs))
    C.fputs(cs, (*C.FILE)(C.stdout))
}
func main() {
	Print("abc")
}