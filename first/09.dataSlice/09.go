package main

import "fmt"

func main() {
	var bulan = [...]string{
		"januari",
		"februari",
		"Maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = bulan[4:7]
	fmt.Println(slice1)      // [mei juni juli]
	fmt.Println(len(slice1)) // 3
	fmt.Println(cap(slice1)) // 8 = capacity mei ke desember

	bulan[5] = "dibuah"
	fmt.Println(slice1) //[mei dibuah juli]

	slice1[0] = "mei lagi"
	fmt.Println(bulan) // [januari februari Maret april mei lagi dibuah juli agustus september oktober november desember]

	// var slice2 = bulan[11:]
	var slice2 = bulan[2:4]
	fmt.Println(slice2) // [desember]

	// append, jika sudah melebihi capacity, maka array baru
	var slice3 = append(slice2, "eko")
	fmt.Println(slice3) // [desember eko]
	slice3[1] = "bukan desember"

	fmt.Println(slice3) // [desember bukan desember]

	fmt.Println(slice2) // [desember]
	fmt.Println(bulan)  // [januari februari Maret april mei lagi dibuah juli agustus september oktober november desember]

	newSlice := make([]string, 2, 5)
	newSlice[0] = "eko"
	newSlice[1] = "kurni"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [...]int{1, 3, 5, 7}
	iniSlice := []int{4, 34, 5, 7} // tanpa angka atau ... type slice

	fmt.Println(cap(iniArray))
	fmt.Println(cap(iniSlice))
}
