package main

import (
	f "fmt"
	g "flag"
)

func main() {
	var name string
	g.StringVar(&name, "nm", "a fake person", "this is the name")
	g.Parse()
	f.Println("hello "+name + "!")
}