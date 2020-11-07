package main

import (
	"fmt"
	"math"
)

func main() {
	// Round -> membulatkan ke atas atau bawah sesuai angka yg paling dekat
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.3))

	// Floor -> mamaksa membulatkan ke bawah
	fmt.Println(math.Floor(1.7))
	// Floor -> mamaksa membulatkan ke atas
	fmt.Println(math.Ceil(1.3))

	// Max -> mengechek yang paling besar
	// Min -> mengechek yang paling kecil
	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))
}
