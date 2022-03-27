package main

import (
	"fmt"
)

func main() {
	var x bool = true
	fmt.Printf("value is %v\t type is %T\n", x, x)

	n := 1
	m := 2
	fmt.Printf("value is %v\t type is %T\n", n, n)
	fmt.Printf("value is %v\t type is %T\n", m, m)

	fmt.Printf("m is equals m %v\n", n == m)

	var dfBool bool
	fmt.Printf("value is %v\t type is %T\n", dfBool, dfBool)

	num := 1
	fmt.Printf("value is %v\t type is %T\n", num, num)

	var num2 uint16 = 1
	fmt.Printf("value is %v\t type is %T\n", num2, num2)

	// Arithmetic operations
	a := 10
	b := 3
	fmt.Println("a+b=", a+b)
	fmt.Println("a-b=", a-b)
	fmt.Println("a*b=", a*b)
	fmt.Println("a/b=", a/b)
	fmt.Println("a mod b=", a%b)
	fmt.Println("a^b=", a^b)
	fmt.Println("a<<b=", a<<b)
	fmt.Println("a>>b=", a>>b)
	fmt.Println("a&b=", a&b)
	fmt.Println("a|b=", a|b)
	fmt.Println("a==b=", a == b)
	fmt.Println("a!=b=", a != b)
	fmt.Println("a>b=", a > b)
	fmt.Println("a<b=", a < b)
	fmt.Println("a>=b=", a >= b)
	fmt.Println("a<=b=", a <= b)
	fmt.Println("a&^b=", a&^b)
	fmt.Println("a^b^b=", a^b^b)
	fmt.Println("a<<b<<b=", a<<b<<b)
	fmt.Println("a>>b>>b=", a>>b>>b)
	fmt.Println("a&b&b=", a&b&b)
	fmt.Println("a|b|b=", a|b|b)

	flotingPont := 3.14
	flotingPont = 13.1e14
	flotingPont = 13.1e-14
	fmt.Printf("value is %v\t type is %T\n", flotingPont, flotingPont)

	var complexNum complex64 = 1 + 2i
	fmt.Printf("value is %v\t type is %T\n", complexNum, complexNum)

	fmt.Printf("Complex number: %v type of complex number %T  \nReal part %v and type of Real Part %T \nImaginary part %v type of  Imaginary part %T \n", complexNum, complexNum, real(complexNum), real(complexNum), imag(complexNum), imag(complexNum))
	fmt.Println("Complex number:", complex(1, 2))

	// operations
	complexNum1 := 1 + 2i
	complexNum2 := 6i
	fmt.Printf("value is %v\t type is %T\n", complexNum1, complexNum1)
	fmt.Printf("value is %v\t type is %T\n", complexNum2, complexNum2)

	fmt.Println("complexNum1+complexNum2=", complexNum1+complexNum2)
	fmt.Println("complexNum1-complexNum2=", complexNum1-complexNum2)
	fmt.Println("complexNum1*complexNum2=", complexNum1*complexNum2)
	fmt.Println("complexNum1/complexNum2=", complexNum1/complexNum2)

	//String

	str := "This is a string"
	fmt.Printf("value is %v\t type is %T\n", str, str)
	fmt.Printf("value is %v\t type is %T\n", str[5], str[5])
	fmt.Printf("value is %c\t type is %T\n", str[5], str[5])

	c := 'a'
	fmt.Printf("value is %v\t type is %T\n", c, c)
	var chr rune = 'a'
	fmt.Printf("value is %v\t type is %T\n", chr, chr)
}
