package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to class on structs")
	abin := User{"Abin Joseph", "itsmeabjs@gmail.com", 23, 8921372134, true}
	fmt.Println(abin)
	fmt.Printf("Abin detailers are %+v \n", abin)

}

type User struct {
	Name   string
	Email  string
	Age    int
	Phone  int
	Status bool
}
