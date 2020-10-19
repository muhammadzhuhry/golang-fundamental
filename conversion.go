package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	// will print -96, it called integer overflow when it touch the boundary it will back to the minimum value
	fmt.Println(nilai8)

	var name = "Zhuhry"
	z := name[0]
	fmt.Println(z, string(z))
}
