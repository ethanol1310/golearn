package main

import "fmt"

func main() {
	// Default
	var num0 int

	// Initialize value
	var num1 int = 1

	// Pass type
	var num2 = 20
	fmt.Println(num0, num1, num2)

	// Short variable declaration
	num := 30
	// only for new variables
	// no new variables on left side of :=
	// num := 31

	num += 1
	fmt.Println("+=", num)

	num++
	fmt.Println("++", num)

	// camelCase - accepted style
	userIndex := 10
	fmt.Println(userIndex)

	// multiple variable declaration
	var weight, height int = 10, 20

	// assignment to existing variables
	weight, height = 11, 21

	// short assignment
	// at least one variable must be new!
	weight, age := 12
	fmt.Println(weight, height, age)

}
