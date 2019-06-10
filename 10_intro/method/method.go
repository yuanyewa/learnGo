package main

import "fmt"

type square struct {
	length float32
	width  float32
}

type triangle struct {
	base   float32
	height float32
}

func (s square) area() float32 {
	return s.length * s.width
}

func (t triangle) area() float32 {
	return t.base * t.height / 2
}
func main() {
	s := square{3, 4}
	var t triangle
	t.base = 10
	t.height = 5
	fmt.Println(s.area(), t.area())
}
