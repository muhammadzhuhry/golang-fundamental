package main

import (
	"fmt"
	"reflect"
)

//Parameter yg berada di posisi terkahir pada sebuah function, memiliki kemampuan dijadikan sebuah varArgs(variable argumen)
//Variadic function adalah function yg memiliki paramater varArgs
//parameter varArgs merupakah parameter yg bisa menirima lebih dari satu input, atau bisa di bilang seperti array / slice
//perbedaan varArgas degnan array?
// -> jika parameter bertipe array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
// -> jika parameter menggunakan varArgs, kita bisalangsung mengirim datanya bila lebih dari satu maka cukup dipisah dengan koma

func sumAll(numbers ...int) int {
	total := 0
	fmt.Println(reflect.TypeOf(numbers))
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(10, 80, 10, 40)
	fmt.Println(total)

	//contoh apabila mengirimkan slice ke varArgs
	slice := []int{90, 20, 30, 80}
	total2 := sumAll(slice...)
	fmt.Println(total2)
}
