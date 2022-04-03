package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter User Age: ")
	// comma ok // error ok

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Your age is ", input)
	fmt.Printf("Type of input is %T", input)
	var currantYear int64 = 2022
	age, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Type of input is %T\n", age)
	fmt.Println("Your barth year is ", currantYear-age)

}
