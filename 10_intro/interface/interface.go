package main

import "fmt"

type shape interface {
	area() float32
}

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
	shapes := []shape{square{4, 5}, triangle{6, 7}}
	for _, v := range shapes {
		fmt.Println(v.area())
	}
}
