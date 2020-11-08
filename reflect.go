package main

import (
	"fmt"
	"reflect"
)

//  Salah satu fungsi reflect ialah dpt melihat struktur kode kita pd saat aplikasi sedang berjalan / tercompile
// Reflection sgt berguna ketika kita ingin membuat library yg general sehingga mudah digunakan

type Sample struct {
	Name string `required:"true" max:"10"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()

			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"Eko"}

	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	fmt.Println(IsValid(sample))
	sample.Name = ""
	fmt.Println(IsValid(sample))
}
