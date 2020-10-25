package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	// for with statement
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	slice := []string{"Muhammad", "Athallah", "Zhuhry"}

	for x := 0; x < len(slice); x++ {
		fmt.Println(slice[x])
	}

	// for range -> short way untuk tipe data collection (array, slice, map)
	for index, value := range slice {
		fmt.Println("index", index, "=", value)
	}

	//developer := make(map[string]string)
	developer := map[string]string{}
	developer["name"] = "Muhammad Athallah Zhuhry"
	developer["role"] = "Backend Developer"

	for key, value := range developer {
		fmt.Println(key, "=", value)
	}
}
