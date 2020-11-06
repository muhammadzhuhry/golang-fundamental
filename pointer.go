package main

import "fmt"

// Secara default di golang semua variable itu pass by value, bukan by reference
// Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yg dikirim adalah duplikasi value nya

// Pointer adalah kemampuan membuat reference ke lokasi data di memory yg sama, tanpa menduplikasi data yg sudah ada
// Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference

type Address struct {
	City, Province, Country string
}

func main() {
	// PASS BY VALUE
	//address1 := Address{"Banda Aceh", "Daerah Istimewa Aceh", "indonesia"}
	//address2 := address1
	//
	//address2.City = "Lhoksemawe"
	//
	//fmt.Println(address1)
	//fmt.Println(address2)

	// PAS BY REFERENCE(WITH POINTER)
	address1 := Address{"Banda Aceh", "Daerah Istimewa Aceh", "indonesia"}
	address2 := &address1 // tambahkan operator & untuk mereference ke alamat memori variabel yg dituju
	var address3 *Address = &address1

	address2.City = "Lhoksemawe" // jika sudah mereference memori jika address2 di update maka address1 juga ikut terupdate karena memiliki alamat memori yg sama

	// Membuat alamat memory baru hasil reference dari memory awal
	//address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}

	// Merubah alamat semua variable yg mereference ke memory awal, maka address1 dan 2 akan memiliki value yg sama karna yg diupdate semua field di alamat memory awal
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(*address2)
	fmt.Println(*address3)

	// Mengunakan function new utk membuat pointer

	// cara biasa membuat pointer
	//var address4 *Address = &Address{"Surabaya", "Jawa Timur", "Indonesia"}

	// menggunakan new
	var address4 *Address = new(Address) // dengan menggunakan new maka ini akan bernilai pointer kosong

	fmt.Println(address4)
}
