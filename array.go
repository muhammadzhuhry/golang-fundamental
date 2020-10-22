package main

import "fmt"

//Array harus berisikan kumpulan data dengan tipe yg sama
//Ketika membuat array, jumlah data yang bisa ditampung harus ditentukan
//Daya tampung Array tidak bisa bertambah setelah Array dibuat

func main() {
	var names [3]string

	names[0] = "Muhammad"
	names[1] = "Athallah"
	names[2] = "Zhuhry"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{80, 90, 100}

	fmt.Println((values))
}

//Function Array
//len(arrayName) -> mendapatkan panjang array
//array[index] -> mendapatkan data di posisi index yg diinputkan
//array[index] = value -> merubah nilai array sesuai index yg diinputkan
