package main

import "fmt"

func main() {
	// var fruits = []string{"apple", "grape", "banana", "melon"}
	// fmt.Println(fruits[0]) // "apple

	// var fruitsA = []string{"apple", "grape"}     // ini slice
	// var fruitsB = [2]string{"banana", "melon"}   // ini array
	// var fruitsC = [...]string{"papaya", "grape"} // ini array

	// var newFruits = fruits[0:2]
	// fmt.Println(newFruits) // [apple grape], semua elemen mulai indeks ke-0, hingga sebelum indeks ke-2
	/*
		fruits[0:4] ==> [apple, grape, banana, melon], 	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-4
		fruits[0:0] ==> [],  menghasilkan slice kosong, karena tidak ada elemen sebelum indeks ke-0
		fruits[4:4] ==> [], menghasilkan slice kosong, karena tidak ada elemen yang dimulai dari indeks ke-4
		fruits[4:0] ==> [], error, pada penulisan fruits[a,b] nilai a harus lebih besar atau sama dengan b
		fruits[:] ==> [apple, grape, banana, melon], ssemua elemen
		fruits[2:] ==> [banana, melon], semua elemen mulai indeks ke-2
		fruits[:2] ==> [apple, grape], semua elemen hingga sebelum indeks ke-2
	*/

	var fruits = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	/*
		var fruits = []string{"apple", "grape", "banana", "melon"}
		fmt.Println(len(fruits))  // len: 4
		fmt.Println(cap(fruits))  // cap: 4

		var aFruits = fruits[0:3]
		fmt.Println(len(aFruits)) // len: 3
		fmt.Println(cap(aFruits)) // cap: 4

		var bFruits = fruits[1:4]
		fmt.Println(len(bFruits)) // len: 3
		fmt.Println(cap(bFruits)) // cap: 3
	*/

	var buahs = []string{"apple", "grape", "banana"}
	var newBuahs = append(buahs, "Orange")

	fmt.Println(newBuahs) // [apple grape banana Orange]

	/*
		var fruits = []string{"apple", "grape", "banana"}
		var bFruits = fruits[0:2]

		fmt.Println(cap(bFruits)) // 3
		fmt.Println(len(bFruits)) // 2

		fmt.Println(fruits)  // ["apple", "grape", "banana"]
		fmt.Println(bFruits) // ["apple", "grape"]

		var cFruits = append(bFruits, "papaya")

		fmt.Println(fruits)  // ["apple", "grape", "papaya"]
		fmt.Println(bFruits) // ["apple", "grape"]
		fmt.Println(cFruits) // ["apple", "grape", "papaya"]
	*/

	// copy
	// dst := make([]string, 3)
	// src := []string{"watermelon", "pinnaple", "apple", "orange"}
	// n := copy(dst, src)

	// dst := make([]string, 3)
	// src := []string{"watermelon", "pinnaple", "aple", "orange"}
	// n := copy(dst, src)

	// fmt.Println(dst) // watermelon pinnaple apple
	// fmt.Println(src) // watermelon pinnaple apple orange
	// fmt.Println(n)   // 3

	dst := []string{"potato", "potato", "potato"}
	src := []string{"watermelon", "pinnaple"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n)   // 2 // value minimum len src or len dst

	/*
			var fruits = []string{"apple", "grape", "banana"}
		var aFruits = fruits[0:2]
		var bFruits = fruits[0:2:2]

		fmt.Println(fruits)      // ["apple", "grape", "banana"]
		fmt.Println(len(fruits)) // len: 3
		fmt.Println(cap(fruits)) // cap: 3

		fmt.Println(aFruits)      // ["apple", "grape"]
		fmt.Println(len(aFruits)) // len: 2
		fmt.Println(cap(aFruits)) // cap: 3

		fmt.Println(bFruits)      // ["apple", "grape"]
		fmt.Println(len(bFruits)) // len: 2
		fmt.Println(cap(bFruits)) // cap: 2
	*/

}
