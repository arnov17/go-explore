package main

import "fmt"

func main() {
	var name [3]string // angka 3 u/ batas maksimal length isi arr

	name[0] = "arnov"
	name[1] = "julian"
	name[2] = "arr"
	// name[3] = "jika" // error out of bounds
	
	fmt.Println(name)

	var values = [3]int {
		93, 83, 88,
	}

	fmt.Println(values)
	fmt.Println(values[2]) // 88 index ke-2

	var lagi[10] int
	fmt.Println(lagi) // hasilnya 10, mesikpun belum ada isi value
	fmt.Println(len(name))

}