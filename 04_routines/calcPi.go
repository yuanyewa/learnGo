package main

import (
	"fmt"
	"math"
)

func pi(ch chan float64, k float64) {
	ch <- -4 * math.Pow(-1, k) / (2 * k * (2*k + 1) * (2*k + 2))
}

func main() {
	var v float64 = 3
	i := 100000
	ch := make(chan float64)
	for k := 1; k < i; k++ {
		go pi(ch, float64(k))
	}
	for k := 1; k < i; k++ {
		v += <-ch
		if k%100 == 0 {
			fmt.Println(v)
		}
	}
}
