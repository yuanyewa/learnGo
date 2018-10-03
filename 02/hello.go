package main

import (
	f "fmt"
	g "flag"
	"os"
	"strconv"
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
}