package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Empty string by default
	var str string

	// With special characters
	var hello string = "Hello\n\t"

	// No special characters
	var world string = "World"

	fmt.Println("str", str)
	fmt.Println("hello", hello)
	fmt.Println("world", world)

	// UTF-8 out of the box
	var helloWorld = "Hello World!"
	hi := "你好 ， 世界"
	fmt.Println("helloWorld", helloWorld)
	fmt.Println("hi", hi)

	// single quotes for bytes (uint8)
	var rawBinary byte = '\x27'

	// rune (uint32) for UTF-8 characters
	var someChinese rune = '茶'

	fmt.Println(rawBinary, someChinese)

	helloWorld = "Hello World, Xin chào"
	// string concatenation
	andGoodMorning := helloWorld + " and good morning!"

	fmt.Println(helloWorld, andGoodMorning)

	// strings are immutatble
	// canot assign to helloWorld[0]
	// helloWorld[0] = 72

	// get string length
	byteLen := len(helloWorld)                    // 19 bytes
	symbols := utf8.RuneCountInString(helloWorld) // 10 runes

	fmt.Println(byteLen, symbols)

	// get a substring in bytes, not characters
	hello = helloWorld[:21] // Hello, 0-11 bytes
	H := helloWorld[0]      // byte, 72, not "P"
	fmt.Println(hello, H)

	// convert to slice bytes and back
	byteString := []byte(helloWorld)
	helloWorld = string(byteString)
	fmt.Println(byteString, helloWorld)
}
