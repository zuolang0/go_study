package main

/*go 语言实现队列
当你在函数内对参数进行的修改需要保留的时候你需要传递指针
*/
import (
	"fmt"
)

type duilie struct {
	arr1 map[int]string
	arr2 map[int]string
}

func (a *duilie) push(tmp string) {
	var tmp_len = len(a.arr1)
	if tmp_len == 0 {
		a.arr1 = make(map[int]string)
	}
	a.arr1[tmp_len] = tmp
}

func (b *duilie) pop() string {
	var len_arr2 = len(b.arr2)
	if len_arr2 > 0 {
		var tmp_value = b.arr2[len_arr2]
		delete(b.arr2, len_arr2)

		return tmp_value
	} else {
		b.arr2 = make(map[int]string)
		var len_arr1 = len(b.arr1)
		if len_arr1 > 1 {
			for i := 0; i < len_arr1; i++ {
				b.arr2[len_arr1-i] = b.arr1[i]
				delete(b.arr1, i)

			}
			//fmt.Print(b)
			return b.pop()
		} else {
			if len_arr1 == 1 {
				var tmp = b.arr1[0]
				delete(b.arr1, 0)
				return tmp
			} else {
				return "null"
			}
		}
	}
}
func main() {
	var test duilie
	test.push("dddd")
	test.push("cccc")
	//fmt.Print(test)
	fmt.Printf(test.pop())
	//fmt.Print(test)
	test.push("aaaaa")
	test.push("bbbb")

	fmt.Printf(test.pop())
	fmt.Printf(test.pop())
	fmt.Printf(test.pop())
}
