package main

import "fmt"

func main() {
	var name = "Eko"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hai, Boleh kenalan?")
	}

	// if short statement
	if lengthName := len(name); lengthName > 5 {
		fmt.Println("Nama terlalu panjan")
	} else {
		fmt.Println("Nama sudah sesuai")
	}
}
