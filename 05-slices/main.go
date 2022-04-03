package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to class on slices")
	var numberList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numberList)
	fmt.Printf("Type and value: %T %v", numberList, numberList)
	fmt.Println("Length of the list:", len(numberList))

	var fruitList = []string{"Apple", "Orange", "Banana", "Grape", "Pineapple"}
	fmt.Println(fruitList)
	fmt.Printf("Type and value: %T %v\n", fruitList, fruitList)
	fmt.Println("Length of the list:", len(fruitList))
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)
	fmt.Println("Length of the list:", len(fruitList))

	highScores := make([]int, 4)
	highScores[0] = 100
	highScores[1] = 90
	highScores[2] = 80
	highScores[3] = 70
	fmt.Println(highScores)
	highScores = append(highScores, 660, 5656)
	fmt.Println(highScores)
	fmt.Println("Length of the list:", len(highScores))
	fmt.Println(highScores)
	fmt.Println("is sorted ", sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("is sorted ", sort.IntsAreSorted(highScores))

	// how to remove value from slice using index
	var courses = []string{"Go", "Python", "Java", "C++", "C#", "Ruby", "JavaScript"}
	fmt.Println(courses)
	var index int = 2
	fmt.Println("Index of the element to be removed:", index)
	fmt.Println("Value of the element to be removed:", courses[index])
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
