package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var names = []string{"john", "wick"}
	printMessage("hallo", names)

	rand.Seed(time.Now().Unix())
	var randomVAlue int

	randomVAlue = randomWithRange(2, 10)
	fmt.Println("random number:", randomVAlue)

	randomVAlue = randomWithRange(2, 10)
	fmt.Println("random number:", randomVAlue)

	randomVAlue = randomWithRange(2, 10)
	fmt.Println("random number:", randomVAlue)

	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)

	var diameter float64 = 15
	var area, circumference = calculated(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divider by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)

}

// func calculated(d float64) (float64, float64) {
// 	// hitung luas lingkaran
// 	var area = math.Pi * math.Pow(d/2, 2)
// 	// keliling
// 	var circumference = math.Pi * d
// 	// mengembalikan 2 nilai
// 	return area, circumference
// }

func calculated(d float64) (area float64, circumference float64) {
	// hitung luas lingkaran
	area = math.Pi * math.Pow(d/2, 2)
	// keliling
	circumference = math.Pi * d
	// mengembalikan 2 nilai
	return
}
