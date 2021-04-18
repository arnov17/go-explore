package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string // bentuk panjang
	person := map[string]string{
		"name":    "arnov",
		"address": "depok",
	}

	fmt.Println(person)            // map[address:depok name:arnov]
	fmt.Println(person["name"])    // arnov
	fmt.Println(person["address"]) // depok

	person["title"] = "programmer"
	fmt.Println(person) // map[address:depok name:arnov title:programmer]

	// var book map[string]string = make(map[string]string)
	book := make(map[string]string)
	book["title"] = "belajar code"
	book["author"] = "acvv"
	book["ups"] = "salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))

}
