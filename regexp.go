package main

import (
	"fmt"
	"regexp"
)

// Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
// Regular expressino di Go-Lang menggunakan library C yg dibuat Google bernama RE2
// https://github.com/google/re2/wiki/Syntax

func main() {
	regex := regexp.MustCompile("e([a-z])o")
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eDo"))

	search := regex.FindAllString("eko eka edo eto eyo eki", 3)
	fmt.Println(search)
}
