package main

import "fmt"

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("angka", i)
	// }

	// var i = 0
	// for i < 5 {
	// 	fmt.Println("angka", i)
	// 	i++
	// } // konsepnya sama seperti while js

	// var i = 0
	// for {
	// 	fmt.Println("Angka", i)

	// 	i++
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("angka", i)
	// }

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop // looping i nya akan berhenti, jika tanpa outerloop, maka tetap melanjutkan increament selanjutny
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}