package main

import "fmt"

func main() {
	for {
		fmt.Println("loop iteration")
		break
	}

	isRun := true
	for isRun {
		fmt.Println("loop iteration with condition")
		isRun = false
	}

	for i := 0; i < 2; i++ {
		fmt.Println("loop iteration", i)
		if i == 1 {
			continue
		}
	}

	s1 := []int{1, 2, 3}
	idx := 0

	for idx < len(s1) {
		fmt.Println("while-style loop, idx", idx, "value", s1[idx])
		idx++
	}

	for i := 0; i < len(s1); i++ {
		fmt.Println("c-style loop", i, s1[i])
	}

	for idx := range s1 {
		fmt.Println("range slice by index", s1[idx])
	}

	for idx, val := range s1 {
		fmt.Println("range slice by idx-value", idx, val)
	}

	profile := map[int]string{
		1: "Vasily",
		2: "Romanov",
	}

	for key := range profile {
		fmt.Println("range map by key", key)
	}

	for _, val := range profile {
		fmt.Println("range map by val", val)
	}

	str := "Xin chào!"
	for pos, char := range str {
		fmt.Printf("%#U at pos %d\n", char, pos)
	}
}
