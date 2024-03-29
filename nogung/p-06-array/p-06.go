package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "trafa"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"
	// names[4] = "ez" // baris kode ini menghasilkan error arena melibih length yang sudah diterisi

	fmt.Println(names[0], names[1], names[2], names[3])

	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	var buahs = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(buahs); i++ {
		fmt.Printf("elemen %d : %s \n", i, buahs[i])
	}

	var fruits3 = [4]string{"apple", "grape", "banana", "melon"}

	// for i, fruit := range fruits3 {
	// 	fmt.Printf("elemen %d : %s\n", i, fruit)
	// }

	for _, fruit := range fruits3 {
		fmt.Printf("elemen %s\n", fruit)
	}

	var animal = make([]string, 2)
	animal[0] = "cat"
	animal[1] = "dog"
	// animal[2] = "cow" // error: index out of range [2] with length 2

	fmt.Println(animal)
}
