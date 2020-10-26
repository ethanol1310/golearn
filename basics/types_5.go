package main

type UserID int

func main() {
	idx := 1
	var uid UserID = 42

	// even if the base type is the same,
	// different types are incompatible
	myID := UserID(idx)

	println(uid, myID)
}
