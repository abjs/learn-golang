package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "Gopher"
	age := 23
	fmt.Printf("Hi i am %s and i am %d Yeses old\n", name, age)

	var ageFloat float32 = float32(age)
	fmt.Printf("Hi i am %s and i am %f Yeses old\n", name, ageFloat)

	var ageString string = strconv.Itoa(age)
	fmt.Printf("Hi i am %s and i am %s Yeses old\n", name, ageString)
}
