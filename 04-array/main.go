package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on array")
	var numberList [10]string
	numberList[0] = "One"
	numberList[1] = "Two"
	numberList[2] = "Three"
	numberList[3] = "Four"
	numberList[4] = "Five"
	numberList[5] = "Six"
	numberList[6] = "Seven"
	numberList[7] = "Eight"
	numberList[8] = "Nine"
	numberList[9] = "Ten"
	fmt.Println(numberList)
	fmt.Println(len(numberList))
	var romanNumbers = [3]string{"I", "II", "III"}
	fmt.Println(romanNumbers)
	fmt.Println(len(romanNumbers))
}
