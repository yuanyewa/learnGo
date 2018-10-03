package main

import (
	f "fmt"
	g "flag"
	"os"
	"strconv"
	"regexp"
)


func main() {
	var name string
	g.StringVar(&name, "nm", "a fake person", "this is the name")
	g.Parse()
	f.Println("hello "+name + "!")

	for _, i := range os.Args {
		f.Println(i)
	}
	f.Println(func() int {x, _ := strconv.Atoi(os.Args[1]); return x}() +1)

	r := regexp.MustCompile(`^@?\w{1,15}$`)
	f.Println(r.MatchString(os.Args[2]))
	f.Println(tmp([]int{100,200,300}...))

	var ar = []int{1,2,3,4,5}
	changeArray(ar)
	f.Println(ar, len(ar), cap(ar))
	{	xx := "a"
		f.Printf("%5s,",xx)
	}
	xx := 3
	f.Println(xx)

}

func changeArray(a []int) { // since slice is a reference, value is changed without needing pointer
	a[0] = 1000
}

func tmp(x ...int) int {
	var y int
	f.Println(x)
	for _, i := range x {
		y += i
	}
	return y
}