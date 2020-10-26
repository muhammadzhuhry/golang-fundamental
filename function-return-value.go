package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hallo, salam kenal!"
	} else {
		return "Hallo " + name
	}
}

func main() {
	fmt.Println(getHello(""))

	result := getHello("Eko")
	fmt.Println(result)
}
