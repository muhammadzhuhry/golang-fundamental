package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpEko NoKTP = "11090998A"
	var marriedStatus Married = true
	fmt.Println(noKtpEko, marriedStatus)
}

// integer

// uint8 -> Unsigned 8-bit integers (0 to 255)
// uint16-> Unsigned 16-bit integers (0 to 65535)
// uint32-> Unsigned 32-bit integers (0 to 4294967295)
// uint64-> Unsigned 64-bit integers (0 to 18446744073709551615)
// int8 -> Unsigned 8-bit integers (-128 to 127)
// int16-> Unsigned 16-bit integers (-32768 to 32767)
// int32-> Unsigned 32-bit integers (-2147483648 to 2147483647)
// int64-> Unsigned 64-bit integers (-9223372036854775808 to 9223372036854775807)

// other

// byte -> uint8
// rune -> int32
// uint -> uint32 or 64 bits(depend on computer)
// int -> same as uint
