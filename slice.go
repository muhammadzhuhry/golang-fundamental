package main

import (
	"fmt"
	"reflect"
)

//Tipe data Slice adalah potongan dari data Array
//Slice mirip dengan Array, yang membedakan adalah ukuran Slice bisa berubah
//Slice dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian atau seluruh data di Array

//Tipe data Slice memiliki 3 data, yaitu pointer, length dan capacity
//Pointer adalah penunjuk data pertama di Array pada Slice
//Length adalah panjang dari Slice
//Capacity adalah kapasitas dari Slice, dimana Length tidak boleh lebih dari Capacity

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // -> 3 : panjang slice dari index ke 4(Mei) sampai index ke 7 - 1(Juli)
	fmt.Println(cap(slice1)) // -> 8 : kapasitasnya dihitung dari index ke 4 sampai index terakhir dari array

	//fmt.Println(slice1)
	//slice1[1] = "Bulan Juni" // -> ini merubah nilai asli di array
	//fmt.Println(slice1)
	//fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2) // -> [November, Desember] : [10:] artinya mengambil nilai dari index ke 10 sampai index paling akhir

	fmt.Println(cap(slice2))
	var slice3 = append(slice2, "Bulan13")
	fmt.Println(slice3)

	slice3[1] = "Bukan Desember" // -> ini tidak merubah nilai array dan slice2 karena kapasitas slice2 adalah 2 jadi ketika di append(menambahkan data) karena kapasitas sudah penuh maka ini membentuk array baru
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Muhammad"
	newSlice[1] = "Zuhri"

	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray, iniSlice, reflect.TypeOf(iniArray), reflect.TypeOf(iniSlice))
}

//Function Slice
//len(sliceName) -> mendapatkan panjang slice
//cap(sliceName) -> mendapatkan kapasitas
//append(slice, data) -> membuat slice baru dengan menambahkan data ke posisi terakhir slice, jika kapasitas sudah penuh maka akan membuat array baru
//make([]TypeData, length, capacity) -> membuat slice baru
//copy(destination, source) -> menyalin slice dari source ke destination

//  Array vs Slice
// iniArray := [...]int{1,2,3,4,5} -> bila memberikan parameter di dalam kurawal(... atau int) maka array
// iniSlice := []int{1,2,3,4,5} -> bila tidak memberikan parameter di kurawal maka slice
