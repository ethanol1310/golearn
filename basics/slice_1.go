package main

import "fmt"

func main() {
	//
	var buf0 []int             // len=0, cap=0
	buf1 := []int{}            // len=0, cap=0
	buf2 := []int{42}          // len=1, cap=1
	buf3 := make([]int, 0)     // len=0, cap=0
	buf4 := make([]int, 5)     // len=5, cap=5
	buf5 := make([]int, 5, 10) // len=5, cap=10

	println(buf0, buf1, buf2, buf3, buf4, buf5)

	someInt := buf2[0]

	fmt.Println(someInt)

	var buf []int
	buf = append(buf, 9, 10)
	buf = append(buf, 12)

	otherBuf := make([]int, 3)
	buf = append(buf, otherBuf...)

	fmt.Println(buf, otherBuf)

	var bufLen, bufCap int = len(buf), cap(buf)

	fmt.Println(bufLen, bufCap)
}
