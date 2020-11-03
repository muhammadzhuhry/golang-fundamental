package main

import "fmt"

//Recover adalah function yang bisa kita gunakan untuk menangkap data panic
//Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan

func endApps() {
	message := recover()
	if message != nil {
		fmt.Println("error message:", message)
	}
	fmt.Println("end app")
}

func runApps(error bool) {
	defer endApps()
	if error {
		panic("APP ERROR")
	}
	fmt.Println("app running properly")
}

func main() {
	//runApps(false)
	runApps(true)
	fmt.Println("app still running")
}
