package main

import "fmt"

func main() {
	name := "Eko"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Salam kenal")
	}

	// switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch without condition
	value := 70

	switch {
	case value >= 80:
		fmt.Println("Nilai anda A")
	case value >= 60:
		fmt.Println("Nilai anda B")
	default:
		fmt.Println("Nilai anda C")
	}
}
