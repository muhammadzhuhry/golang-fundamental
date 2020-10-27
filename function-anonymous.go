package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

//func blacklistAdmin(name string) bool {
//	return name == "Admin"
//}
//
//func blacklistRoot(name string) bool {
//	return name == "Root"
//}

func main() {
	//registerUser("Admin", blacklistAdmin)

	// contoh anonymous function di variable
	blacklistAdmin := func(name string) bool {
		return name == "Admin"
	}
	registerUser("Admin", blacklistAdmin)

	// contoh anonymous function di parameter
	registerUser("Root", func(name string) bool {
		return name == "Root"
	})
}
