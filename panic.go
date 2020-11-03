package main

import "fmt"

// Panic function adalah function yg bisa kita gunakan untuk menghentikan program
// Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
// Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi

func endApp() {
	fmt.Println("end app")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APP ERROR")
	}
	fmt.Println("app running properly")
}

func main() {
	//runApp(false)
	runApp(true)
}
