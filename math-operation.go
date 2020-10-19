package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a, b = 10, 10

	c := a * b
	fmt.Println(c)

	// Augmented Assignment
	var i = 10
	i += 10
	fmt.Println(i)

	// unary operator
	i++
	fmt.Println(i)

}
