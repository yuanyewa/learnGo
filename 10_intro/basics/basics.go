package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world!")
	var x int = 1
	x++
	fmt.Println(x)

	var y []float32 = []float32{math.Pi, math.E}
	fmt.Println(y)

	z := complex(1.2, 3.4)
	fmt.Println(z)

	s := "this is a string"
	fmt.Println(s + "   blabal")

	m := map[int]string{
		0: "red",
		1: "green",
		2: "blue",
	}
	if v, ok := m[2]; ok {
		fmt.Println(v)
	}
	for i, v := range s {
		fmt.Println(i, string(v))
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i * i)
	}
	i := 10
	switch i {
	case 10:
		fmt.Println("I==10 satisfied!")
	default:
		fmt.Println("fall back!")
	}

	y = append(y, 1, 2, 3, 4)
	fmt.Println(y[0:4])

	fmt.Println(min(10, 3))

	x = 100
	incVal(&x)
	fmt.Println(x)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func incVal(x *int) {
	*x = *x + 1
}
