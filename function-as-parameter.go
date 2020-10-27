package main

import "fmt"

//func sayHelloWithFilter(name string, filter func(string) string) {
//	nameFiltered := filter(name)
//	fmt.Println("Hello", nameFiltered)
//}

// example if using function as parameter with type declaration

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Eko", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
