package main

import "fmt"

func main() {
	var name string

	name = "Mohd Zhuhry"
	fmt.Println(name)

	name = "Muhammad Zhuhry"
	fmt.Println((name))

	var friendName = "Budi"
	fmt.Println(friendName)

	var age = 23
	fmt.Println(age)

	country := "indonesia"
	fmt.Println(country)

	var (
		firstName = "Muhammad"
		lastName  = "Zhuhry"
	)
	fmt.Println(firstName, lastName)

	first, last := "budi", "bambang"
	fmt.Println(first, last)
}
