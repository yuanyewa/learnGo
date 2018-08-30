package main

import (
	"fmt"
)

func main2() {
	fmt.Println("this is me")
}

func main() {
	fmt.Println("this is the main func")
	main2()
}
