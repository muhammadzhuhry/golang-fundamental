package main

import "fmt"

func main() {
	var name1 = "Zuhri"
	var name2 = "Zuhri"

	var result bool = name1 == name2
	fmt.Println(result)

	value1, value2 := 100, 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
