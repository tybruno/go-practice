package assignment_2

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

func main() {
	s1 := square{sideLength: 2}
	t1 := triangle{height: 2, base: 2}

	printArea(s1)
	printArea(t1)
}

func (s square) getArea() float64 {
	area := int(s.sideLength) * int(s.sideLength)
	return area
}

func (t triangle) getArea() float64 {
	area := 0.5 * int(t.base) * int(t.height)
	return area
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
