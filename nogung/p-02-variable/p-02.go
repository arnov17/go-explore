package main

import "fmt"

func main() {
	var firstName string = "john"

	// var lastName string
	// lastName = "wick"

	// tanpa var
	lastName := "wick"

	middlename := "lia"

	fmt.Printf("halo john wick!\n")

	fmt.Printf("halo %s %s!\n", firstName, lastName)

	fmt.Println("halo", firstName, lastName+"!")

	fmt.Println("halo", middlename, "!")

	name := new(string)

	fmt.Println(name)  // 0x20818a220
	fmt.Println(*name) // ""

	/*
		var first, second, third string
		first, second, third = "satu", "dua", "tiga"

		var fourth, fifth, sixth string = "empat", "lima", "enam"

		seventh, eight, nine := "tujuh", "delapan", "sembilan"

		// multi variable
		one, isFriday, twoPointTwo, say := 1, true, 2.2,

		// Variabel Underscore _
		_ = "belajar golang"
		name, _ := "john", "wick"

	*/

}
