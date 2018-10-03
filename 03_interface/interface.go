package main

// interface:
// define a function that applies to an interface type; this function makes use of the interface functions
// define the interface type with set of functions used by the top-level function
// define the sub-types that you want the top-level function to apply to
// implement the interface function for each type

import (
	"fmt"
	"strconv"
)

func describe(c car) { // top level function; I want to "describe" the car
	fmt.Println("car " + c.name() + " with color " + c.color() + " is made in " + strconv.Itoa(c.year()))
}

type car interface { // "describe" uses the name, color, year of the car
	name() string
	color() string
	year() int
}

type mycar struct{Year int} // my car type
type urcar struct{Year int} // your car type

func (m mycar) name() string { return "ford" } // implement my car functions
func (m mycar) color() string { return "blue" }
func (m mycar) year() int { return m.Year }
func (u urcar) name() string { return "bmw" } // implement your car functions
func (u urcar) color() string { return "black" }
func (u urcar) year() int { return u.Year }

func main() {
	describe(mycar{100}) // describe my car
	describe(urcar{200}) // describe your car
	
	// empty interface allows list of different types
	z := []interface{}{3, "hello", mycar{100}}
	fmt.Println(z)
}