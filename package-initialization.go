package main

import (
	"fmt"
	"github.com/muhammadzhuhry/belajar-golang-dasar/database"
)

// PACKAGE INITIALIZATION

//import _ "github.com/muhammadzhuhry/belajar-golang-dasar/helper"

// Saat kita membuat package, kita bisa membuat sebuah function yg akan di akses pertama kali ketika package kita di akses
// ini sangat cocok ketika contohnya pacage kita berisi function-function untuk berkomunikasi dgn database, kita membuat  function inisialisasi utk membuat koneksi ke db
// Untuk membuat function yg diakses secara otomatis ketika package diakses, kita cukup membuat function dgn nama init

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}

// BLANK IDENTIFIER

// Kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah satu function yg ada di package
// Secara default, Go-Lang akan komplen ketika ada package yg diimport namun tidak digunakan
// Untuk menangani hal tersebut kita bisa menggunakan blank identifier (_) sblm nama package ketika melakukan import
