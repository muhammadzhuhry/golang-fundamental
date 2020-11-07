package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, _ := strconv.ParseBool("true")
	fmt.Println(boolean)

	// base int 10 -> decimal, 2 -> binary, 8 -> oktal, 16 -> hexadecimal
	number, _ := strconv.ParseInt("1000000", 10, 64)
	fmt.Println(number)

	value := strconv.FormatInt(10000000, 10)
	fmt.Println(value)

	// with Atoi no need to declare base int, default is 10 / decimal
	valueInt, _ := strconv.Atoi("200000")
	fmt.Println(valueInt)
}
