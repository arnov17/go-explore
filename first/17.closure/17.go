package main

import "fmt"

func main() {
	counter := 0
	name := "abi"
	kata := "hallo"

	increment := func() {

		counter++
		name = "doi"
		// counter & nama merubah nilai var diatas
		kata := "Hai"
		// tidak merubah nilai var diluar scope atas

		fmt.Println(counter)
		fmt.Println(name)
		fmt.Println(kata)
		// a := "test"
	}
	// a tidak bisa dpinggil, karena diluar scope dibawah
	// fmt.Println(a)

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(kata)
}
