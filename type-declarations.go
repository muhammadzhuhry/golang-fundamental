package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpEko NoKTP = "11090998A"
	var marriedStatus Married = true
	fmt.Println(noKtpEko, marriedStatus)
}
