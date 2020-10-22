package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 75

	var kkmUjian = ujian >= 80
	var kkmAbsensi = absensi >= 80

	var lulus = kkmUjian && kkmAbsensi
	fmt.Println(lulus)
}
