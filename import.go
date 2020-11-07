package main

import (
	"fmt"
	"github.com/muhammadzhuhry/belajar-golang-dasar/helper"
)

func main() {
	helper.SayHello("Zuhri")
	//helper.sayGoodbye("Zuhri") // error
	fmt.Println(helper.Application)
	//fmt.Println(helper.version) // error
}

// ACCESS MODIFIER
// Di bahasa pemrograman lain, biasanya ada kata kunci yg bisa digunakan untuk menentukan access modifier terhadap suatu function atau variable
// Di Golang untuk menentukan access modifier, cukup menggunakan nama function atau variable(jika diawali dgn huruf besar maka artinya bisa di akses di package lain)
