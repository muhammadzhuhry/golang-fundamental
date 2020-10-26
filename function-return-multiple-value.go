package main

import "fmt"

func getFullName() (string, string, int8) {
	return "Muhammad", "Zhuhry", 17
}

func main() {
	first, last, _ := getFullName()
	fmt.Println(first, last)
}
