package main

import "fmt"

//closures adalah kemampuan sebuah function berinteraksi dgn data data disekitarnya dalam scope yg sama

//func increment(value int) int {
//	return value + 1
//}

func main() {
	//fmt.Println(increment(0))

	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}
