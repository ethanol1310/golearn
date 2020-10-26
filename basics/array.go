package main

import "fmt"

func main() {
	// the size of the array is part of its type
	// initialization with default values
	var a1 [3]int
	fmt.Println("a1 short", a1)
	fmt.Printf("a1 short %v\n", a1)
	fmt.Printf("a1 full %#v\n", a1)

	const size = 2
	var a2 [2 * size]bool
	fmt.Println("a2", a2)

	// size determination when declaring
	a3 := [...]int{1, 2, 3}
	fmt.Println("a2", a3)

	// check at compile or run
	// invalid array index 4 (oob for 3 element array)
	// a3[4] = 12
}
