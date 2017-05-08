package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human
	school string
}
type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi I am %s you can call me %s\n", h.name, h.phone)
}
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}
func (s *Student) SayHi() {
	fmt.Printf("Hi, I am %s, I school at %s. Call me on %s\n", s.name,
		s.school, s.phone)
}
func main() {
	mark := Student{Human{"ZL", 25, "12345"}, "WZ"}
	sam := Employee{Human{"ZC", 25, "123456"}, "Yz"}
	mark.SayHi()
	sam.SayHi()
}
