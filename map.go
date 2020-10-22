package main

import "fmt"

//Map adalah kumpulan key value yang dimana key nya bersifat unik tidak boleh sama
//Tipe data valuenya haruslah bertipe yang sama
//Berbeda dengan Array dan Slice data yang dimasukan ke Map boleh sebanyak banyaknya, dengan catatan keynya harus berbeda
//Bila key nya sama maka otomatis key data sebelumnya akan di replace dgn key data yang baru

func main() {
	person := map[string]string{
		"name":    "Zuhri",
		"address": "Aceh",
	}

	person["job"] = "Developer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	books := make(map[string]string)
	books["title"] = "Golang Programming"
	books["author"] = "Muhammad Zhuhry"
	books["typo"] = "typo here"

	fmt.Println(books)
	delete(books, "typo")
	fmt.Println(books)
}

//Function Map
//len(mapName) -> mendapatkan panjang map
//map[key] -> mendapatkan data/value sesuai dengan key yg diinputkan
//map[key] = value -> merubah/menambahkan nilai map sesuai key value yg diinputkan
//make(map[TypeKey]TypeValue) -> membuat map baru
//delete(map, "key") -> menghapuskan data(key value) sesuai dengan key yg diinputkan
