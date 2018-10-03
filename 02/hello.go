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
}