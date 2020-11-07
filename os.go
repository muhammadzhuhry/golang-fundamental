package main

import (
	"fmt"
	"os"
)

// Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen

func main() {
	args := os.Args
	fmt.Println("Argument :", args)

	name, _ := os.Hostname()
	fmt.Println(name)
}
