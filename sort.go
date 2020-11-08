package main

// Package sort adalah package yang berisikan utilitas untuk proses pengurutan
// Agar data bisa diurutkan, kita harus mengimplementasi kontrak interface sort.Interface

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

// bikin function struct len, less, swap

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Eko", 30},
		{"Joko", 35},
		{"Budi", 27},
		{"Rudi", 12},
	}

	fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println("sorted:", users)
}
