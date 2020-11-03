package main

import "fmt"

// Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
// Defer function akan selalu dieksekusi walaupun terjadi error di function yang di eksekusi

// Ex without Defer func logging cannot be execute because of error on result
//func logging() {
//	fmt.Println("selesai memanggil function")
//}
//
//func runApplication(number int) {
//	fmt.Println("run application")
//
//	result := 10 / number
//	fmt.Println(result)
//
//	logging()
//}
//
//func main() {
//	runApplication(0)
//}
//
//// Ex with Defer func logging executed even though there is an error
func logging() {
	fmt.Println("selesai memanggil function")
}

func runApplication(number int) {
	defer logging()
	fmt.Println("run application")

	result := 10 / number
	fmt.Println(result)
}

func main() {
	runApplication(0)
}
