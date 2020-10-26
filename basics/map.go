package main

import "fmt"

func main() {
	// initialization on creation
	var user map[string]string = map[string]string{
		"name":     "Vasily",
		"lastName": "Romanov",
	}

	// immediate with required capacity
	profile := make(map[string]string, 10)

	// amount of elements
	mapLength := len(user)

	fmt.Printf("%d %+v\n", mapLength, profile)

	// if there is no key, it will return the default
	// value for the type
	mName := user["middleName"]
	fmt.Println("mName", mName)

	// check for key existence
	mName, mNameExist := user["middleName"]
	fmt.Println("mName:", mName, "mNameExist:", mNameExist)

	// empty variable - just check that the key is there
	_, mNameExist2 := user["middleName"]
	fmt.Println("mNameExists2", mNameExist2)

	// remove key
	delete(user, "lastName")
	fmt.Printf("%#v\n", user)
}
