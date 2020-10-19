package main

import "fmt"

func main() {
	const (
		firstName string = "Muhammad"
		lastName         = "Zhuhry"
		value            = 1000
	)

	//it's okay if you not use constant variable, go will not giving error
	//fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
}
