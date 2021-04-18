package main

import "fmt"

func main() {
	name := "arnov"
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("nama terlalu panjang")
	// case false:
	// 	fmt.Println("nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama sudah benar")
	default:
		fmt.Println("nama sudah benar")
	}

}
