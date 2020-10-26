package main

import "fmt"

func main() {
	boolVal := true

	// simple condition
	if boolVal {
		fmt.Println("boolVal is true")
	}

	mapVal := map[string]string{"lastName": "Vasily"}

	// condition with an initialization block
	if keyValue, keyExist := mapVal["name"]; keyExist {
		fmt.Println("name = ", keyValue)
	}

	// get only a sign of the key existence
	if _, keyExist := mapVal["name"]; keyExist {
		fmt.Println("key 'name' exist")
	}

	cond := 1
	if cond == 1 {
		fmt.Println("cond is 1")
	} else if cond == 2 {
		fmt.Println("cond is 2")
	}

	// switch
	strVal := "name"
	switch strVal {
	case "name":
		fallthrough
	case "test", "lastName":
		// some work
	default:
		// some work
	}

	// swith as a replacement for many ifelse
	var val1, val2 = 2, 2
	switch {
	case val1 > 1 || val2 < 11:
		fmt.Println("first block")
	case val2 > 10:
		fmt.Println("second block")
	}

	// exit the loop while inside the switch
Loop:
	for key, val := range mapVal {
		println("switch in loop", key, val)
		switch {
		case key == "firstName":
			break
			println("dont pront this")
		case key == "lastName" && val == "Vasily":
			println("switch - break loop here")
			break Loop
		}
	}
}
