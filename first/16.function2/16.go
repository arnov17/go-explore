package main

import "fmt"

// var declar
type Filter func(string) string
type Blacklist func(string) bool

func sayHellotoFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked")
	} else {
		fmt.Println("welcome")
	}
}

func main() {
	sayHellotoFilter("eko", spamFilter)

	// func as value
	filter := spamFilter
	sayHellotoFilter("eko", filter)

	// anonymous func
	blocklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blocklist)
	registerUser("eko", blocklist)

}
