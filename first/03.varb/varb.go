package main

import "fmt"

func main() {
	// deklar var with data type
	var name string

	name = "Arnov"
	fmt.Println(name)

	name = "Julian"
	fmt.Println(name)

	var friendName = "Ani"
	fmt.Println(friendName)

	var age = 26
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	var (
		firstName = "Eko"
		lastName  = "Khannedy"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// deklar var wthout data type
	var herName = "Julian"
	hisname := "kemosabe"
	fmt.Println(herName)
	fmt.Println(hisname)
}
