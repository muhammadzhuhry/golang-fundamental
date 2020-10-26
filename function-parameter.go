package main

import "fmt"

func sayHelloTo(firstName, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHelloTo("Muhammad", "Zhuhry")

	firstName := "Budi"
	sayHelloTo(firstName, "Khanedy")
}
