package main

import "fmt"

//go:generate stringer -type=day
type day int

//comment
const (
	Monday day = iota //this is a comment
	Tuesday
)

func (i day) String() string {
	return [2]string{"Monday", "Tuesday"}[i]
}
func main() {
	today := Tuesday
	fmt.Println(today)
	fmt.Printf("%s\n", today)
}
