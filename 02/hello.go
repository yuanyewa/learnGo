package main

import (
	f "fmt"
	g "flag"
	"os"
)

func main() {
	var name string
	g.StringVar(&name, "nm", "a fake person", "this is the name")
	g.Parse()
	f.Println("hello "+name + "!")

	for _, i := range os.Args {
		f.Println(i)
	}
}