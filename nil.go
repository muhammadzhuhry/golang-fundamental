package main

import "fmt"

// Biasanya di dalam programming language lain, object yg belum diinisialisasi maka secara otomatis nilainya adalah null atau nil
// Di Golang ketika membuat variable dgn tipe data tertentu, maka secara otomatis akan dibuatkan default valuenya / zero valuenya -> (tidak langsung menjadi nil)
// Namun di Golang ada data nil, yaitu data kosoong
// Nil sendiri hanya bisa digunakan di bbrp tipe data, seperti interface, function, map, slice, pointer atau channel

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := NewMap("")
	if person == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(person)
	}
}
