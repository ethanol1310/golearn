package main

import "fmt"

const pi = 3.141

const (
	hello = "Xin chào"
	e     = 2.718
)

const (
	zero = iota // auto increment
	_           // empty variable, skip iota
	two
	three // = 3
)

const (
	_         = iota             //skip the first value
	KB uint64 = 1 << (10 * iota) // 1 << (10 * 1) = 1024
	MB                           // 1 << (10 * 2) = 1048576
)

const (
	// untyped constant -> auto convert
	year = 2017

	// typed constant
	yearTyped int = 2017
)

func main() {
	var month int32 = 13 // int16
	fmt.Println(month + year)

	// month + yearTyped (mismatched types int32 and int)
	// fmt.Println(month + yearTyped)
}
