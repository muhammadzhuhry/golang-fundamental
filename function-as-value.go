package main

import "fmt"

//function adalah first class citizen -> bisa berdiri sendiri
//function juga merupakan tipe data dan bisa disimpan di dalam vairable

func getGoodBye(name string) string {
	return "Good Bye, " + name
}

func main() {
	sayGoodBye := getGoodBye

	result := sayGoodBye("Zuhri")

	fmt.Println(result)
	fmt.Println(sayGoodBye("Bayu"))
	fmt.Println(getGoodBye("Herlambang"))
}
