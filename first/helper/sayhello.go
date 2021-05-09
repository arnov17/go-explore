package helper

import "fmt"

// acces modfifier.
// huruf besar depan untuk punclik
// huruf kecil, private (tidak bisa diakses keluar)
var Application = "Golang"

var version = 1 // tidak bisa diakses keluar

func SayHello(name string) {
	fmt.Println("Hello", name)
	fmt.Println(version)
}

func sayGoodBye(word string) {
	fmt.Println("selamat", word)
}

// func sayGoodbye tidak bisa diakses keluar
