package helper

import "fmt"

var version = 1             // tidak bisa diakses file lain
var Application = "Go-Lang" // bisa diakses file lain

// karena  diawali huruf besar jadi ini bisa diakses dari file lain
func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}
