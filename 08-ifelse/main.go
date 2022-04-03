package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on structs")
	// mark grater than 24 then pass else failed
	mark := 14
	if mark >= 24 {
		fmt.Println("Passed")
	} else if mark <= 18 {
		fmt.Println("You are are under 18")
	} else {
		fmt.Println("Failed")
	}

}
