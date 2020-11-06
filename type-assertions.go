package main

import "fmt"

// Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
// Fitur ini sering digunakan ketika kita bertemu dengan data interface kosong

func random() interface{} {
	return 1
}

func main() {
	// TYPE ASSERTION
	//var result = random()
	//var resultString = result.(string) // disini type assertions merubah return interface kosong menjadi string
	// pastikan return dr function tersebut sesuai, pada contoh diatas string dan ketika type assertion yg dikirim juga string, apabila beda tipe data maka akan terjadi panic

	//fmt.Println(resultString)

	// ini contoh tipe data yg tidak sesuai antar function dan pemanggilnya maka akan terjadi panic
	//var resultBool = result.(bool)
	//fmt.Println(resultBool)

	// TYPE ASSERTION SWITCH
	// Saat salah menggunakan type assertion (tipe data tidak sesuai antar function dan pemanggilnya), maka bisa berakibat terjadi panic di aplikasi
	// Jika panic dan tidak ter-recover, maka program otomatis akan mati
	//	Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertion

	var result = random()

	switch value := result.(type) {
	case string:
		fmt.Println("value", value, "is string")
	case int:
		fmt.Println("value", value, "is int")
	default:
		fmt.Println("unknown type")
	}
}
