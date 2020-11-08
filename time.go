package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year(), now.YearDay())

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02" // time.RFC339(for layout check this out)
	parse, _ := time.Parse(layout, "2020-03-20")
	fmt.Println(parse)
}
