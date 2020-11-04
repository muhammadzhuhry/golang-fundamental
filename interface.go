package main

import "fmt"

// Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
// Sebuah interface berisikan definisi-definisi method
// Biasanya interface digunakan sebagai kontrak

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func main() {
	//zuhri := Person{Name: ""}
	//zuhri.Name = "ekoo"

	zuhri := Person{
		Name: "Muhammad Athallah Zhuhry",
		Age:  23,
	}
	SayHello(zuhri)

	var biawak Animal
	biawak.Name = "Biawak"
	SayHello(biawak)
}

// Implementasi interface

// setiap tipe data yg sesuai dengan kontrak interface, secara otomatis dianggap sbg interface tersebut
// sehingga kita tidak perlu mengimplementasikan interface secara manual
// hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface mana

type Person struct {
	Name string
	Age  int
}

// create struct method
func (prsn Person) GetName() string {
	return prsn.Name
}

type Animal struct {
	Name string
}

func (anml Animal) GetName() string {
	return anml.Name
}
