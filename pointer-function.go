package main

import "fmt"

// Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function
// Oleh karena itu, jika kita ingin mengubah data di dalam function, data yg asli tidak akan pernah berubah
// Hal ini membuat variable menjadi aman, karena tidak akan bisa dirubah
// Namun ada beberapa case yg kadang mengharuskan kita untuk membuat function yg bisa merubah data asli parameter tersebut
// Untuk melakukan ini, kita juga bisa menggunakan pointer di function
// Untuk menjadikan sebuah paramater sebagai pointer, kita bisa menggunakan operator * di parameternya

type Alamat struct {
	City, Province, Country string
}

func changeCountryToIndonesia(alamat *Alamat) {
	alamat.Country = "Indonesia"
}

func main() {
	adrres := Alamat{"Subang", "Jawa Barat", ""}
	changeCountryToIndonesia(&adrres)

	fmt.Println(adrres)
}
