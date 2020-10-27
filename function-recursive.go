package main

import "fmt"

// Recursive function adalah function yang memanggil dirinya sendiri
// Contoh kasus yang lebih mudah menggunakan recursive function adalah faktorial

func factorialLoop(value int) (result int) {
	result = 1

	for i := value; i > 0; i-- {
		result *= i
	}
	return
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)

	recursive := factorialRecursive(5)
	fmt.Println(recursive)
}
