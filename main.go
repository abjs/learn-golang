package main

import "fmt"

func main() {
	const pi = 3.14
	fmt.Printf("value is %v\t type is %T\n", pi, pi)
	const a int = 10
	const b float32 = 10.26
	const c string = "hello"
	const d bool = true
	fmt.Printf("value is %v\t type is %T\n", a, a)
	fmt.Printf("value is %v\t type is %T\n", b, b)
	fmt.Printf("value is %v\t type is %T\n", c, c)
	fmt.Printf("value is %v\t type is %T\n", d, d)

}
