package main

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte
type Box struct {
	width, height, depth float64
	color                Color
}
type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}
func (b Box) SetColor(c Color) {
	b.color = c
}
func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}
func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}
func (c Color) string() string {
	string := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return string[c]
}
func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}
	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("the volume of the first on is", boxes[0].Volume(), "cm³")
	fmt.Println("the color of the last one is", boxes[len(boxes)-1].color.string())
	fmt.Println("the biggest one is", boxes.BiggestColor().string())
	fmt.Println("let`s paint them all black")
	boxes.PaintItBlack()
	fmt.Println("the color of the second one os", boxes[1].color.string())
	fmt.Println("obviously,now,the biggest one is", boxes.BiggestColor().string())
}
