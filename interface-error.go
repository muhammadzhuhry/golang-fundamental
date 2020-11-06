package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	//var contohError error = errors.New("Upss error")
	//fmt.Println(contohError.Error())

	hasil, err := Pembagi(100, 0)
	// jika tidak ada error
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
