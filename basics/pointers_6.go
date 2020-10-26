package main

import "fmt"

func main() {
	a := 2
	b := &a
	*b = 3  // a = 3
	c := &a // new pointer to variable a

	fmt.Println(c)

	// getting a pointer to a variable of type int
	// initialized with default value
	d := new(int)
	*d = 12
	*c = *d // c = 12 -> a = 12
	*d = 13 // c and a have not changed

	c = d   // now c points where d is
	*c = 14 // c = 14 -> d = 14, a = 12
}
