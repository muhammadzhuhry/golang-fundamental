package main

// Package strings adalah package yang berisikan function function untuk memanipulasi tipe data string

import (
	"fmt"
	"strings"
)

func main() {
	// mengecek apakah string di params1 mengandung string yg di params2
	fmt.Println(strings.Contains("Muhammad Zhuhry", "Zhuhry"))
	fmt.Println(strings.Contains("Muhammad Zhuhry", "Budi"))

	// Split -> memotong string berdasarkan seperator menjadi slice
	fmt.Println(strings.Split("Muhammad Zhuhry", " "))

	// ToLower -> membuat semua karakter string menjadi lower case
	fmt.Println(strings.ToLower("Muhammad Zhuhry"))
	// ToUpper -> membuat semua karakter string menjadi upper case
	fmt.Println(strings.ToUpper("Muhammad Zhuhry"))

	// Trim -> memotong cutset di awal dan akhir string
	fmt.Println(strings.Trim("   Muhammad Athallah Zhuhry   ", " "))

	// ReplaceAll -> mengubah semua string dari from ke to
	fmt.Println(strings.ReplaceAll("Zuhri Zuhri Zuhri Athallah Zuhri", "Athallah", "Budi"))
}
