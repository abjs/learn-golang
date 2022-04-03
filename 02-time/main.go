package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Wellcome to time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-06-2005 15:04:05 Monday"))

	createdData := time.Date(2022, time.April, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdData)
	fmt.Println(createdData.Format("01-06-2005 15:04:05 Monday"))

}
