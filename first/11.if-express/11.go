package main

import "fmt"

func main() {
	// if secara umum sama dengan js
	var name = "arnov"

	// var length = len(name)
	// if(length > 5) {
	// 	fmt.Println("terlalu panjang")
	// } else {
	// 	fmt.Println("nama sudah benar")
	// }

	// short statement if
	if length := len(name); length > 5 {
		fmt.Println("terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}

}
