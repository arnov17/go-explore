package main

import "fmt"

func main() {
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("perulangan ke ", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke", counter)
	}

	slice := []string{"aju", "aku", "ika", "koi"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}
	// tanda _ adalah i yang tidak dibutuhkan, karena jika ada i, harus dipanggil
	for _, value := range slice {
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "eko"
	person["title"] = "akuntan"

	for key, value := range person {
		fmt.Println("key", key, "=", "value", value)
	}
}
