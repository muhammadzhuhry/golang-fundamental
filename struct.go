package main

import "fmt"

//Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
//Struct biasanya representasi data dalam program aplikasi yang kita buat
//Data di struct disimpan di field
//Sederhananya struct adalah kumpulan dari field

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var zuhri Customer

	zuhri.Name = "Muhammad Athallah Zhuhry"
	zuhri.Address = "Indonesia"
	zuhri.Age = 23

	fmt.Println(zuhri)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}

	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 25}

	fmt.Println(budi)

	// called struct method
	zuhri.sayHi("Iqbal")
}

// Struct method / function
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}
