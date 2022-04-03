package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")
	// var ptr *int
	// fmt.Println("The value of ptr is ", ptr) //The value of ptr is  <nil>
	myNumber := 23
	var ptr = &myNumber
	fmt.Println("The value of ptr is ", ptr)                      //The value of ptr is  0xc000014098
	fmt.Printf("The value of ptr is %v and type %T\n", *ptr, ptr) //The value of ptr is  0xc000014098
	*ptr = 6
	fmt.Println("The value of myNumber is ", myNumber) //The value of ptr is  6

}
