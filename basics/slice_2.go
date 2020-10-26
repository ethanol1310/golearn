package main

import "fmt"

func main() {
	buf := []int{1, 2, 3, 4, 5}
	fmt.Println(buf)
	println(buf)

	// get a slice poiting to the same memory
	sl1 := buf[1:4] // [2, 3, 4]
	sl2 := buf[:2]  // [1, 2]
	sl3 := buf[2:]  // [3, 4, 5]
	fmt.Println(sl1, sl2, sl3)
	println(sl1, sl2, sl3)

	newBuf := buf[:] // [1, 2, 3, 4, 5]
	// buf = [9, 2, 3, 4, 5] because the same memory
	newBuf[0] = 9
	println(newBuf)

	// newBuf now points to other data
	newBuf = append(newBuf, 6)
	println(newBuf)

	// buf = [9, 2, 3, 4, 5] unchanged
	// newBuf = [1, 2, 3, 4, 5, 6] changed
	newBuf[0] = 1
	fmt.Println("buf", buf)
	fmt.Println("newBuf", newBuf)

	// Copy one slice to another
	var emptyBuf []int // len=0, cap=0

	// Wrong - copy less (by len) of 2 slices
	copied := copy(emptyBuf, buf) // copied = 0
	fmt.Println(copied, emptyBuf)

	// right
	newBuf = make([]int, len(buf), len(buf))
	copy(newBuf, buf)
	fmt.Println(newBuf)

	// can be copied to part of an existing slice
	ints := []int{1, 2, 3, 4}
	copy(ints[1:3], []int{5, 6})
	fmt.Println(ints)
}
