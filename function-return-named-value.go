package main

import "fmt"

func getFullName2() (firstName, middleName, lastName string) {
	firstName, middleName, lastName = "Muhammad", "Athallah", "Zhuhry"
	return
}

func main() {
	first, middle, last := getFullName2()
	fmt.Println(first, middle, last)
}
