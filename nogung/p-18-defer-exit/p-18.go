package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("halo")
	fmt.Println("selamat datang")

	orderSomeFood("pizza")
	orderSomeFood("burger")

	number := 3

	if number == 3 {
		fmt.Println("halo 1")
		// defer fmt.Println("halo 3") akan tetap halo 1,2,3
		func() {
			defer fmt.Println("halo 3")
		}()
	}

	fmt.Println("halo 2")

	// EXIT akan tidak mengeksusi func ini
	// defer fmt.Println("halo")
	// os.Exit(1)
	// fmt.Println("selamat datang")

}

func orderSomeFood(menu string) {
	defer fmt.Println("Terimakasih, silakan tunggu")
	if menu == "pizza" {
		fmt.Print("Pilihan tepat!", " ")
		fmt.Print("Pizza ditempat kami paling enak!", "\n")
		return
	}

	fmt.Println("Pesanan anda:", menu)
}
