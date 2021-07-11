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

	var angka = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var avg = calculate(angka...)
	var msg = fmt.Sprintf("rata-rata %.2f \n", avg)
	fmt.Print(msg)

	yourHobiies("wick", "sleep", "eating")
	var hobbies = []string{"reading", "sport"}
	yourHobiies("ani", hobbies...)

	// closure
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				max = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numbers, min, max)

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println("original number :", numbers)
	fmt.Println("original number :", newNumbers)

	var newMax = 3
	var howMany, getNumbers = findMax(numbers, newMax)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)    // 9
	fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]

	var data = []string{"wick", "jason", "delta"}
	var dataConstainO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t \t:", data)
	fmt.Println("filter ada data o \t", dataConstainO)
	fmt.Println("data len 5 \t", dataLength5)

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

func calculate(number ...int) float64 {
	var total int = 0
	for _, num := range number {
		total += num
	}

	var avg = float64(total) / float64(len(number))
	return avg
}

func yourHobiies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("hello, My name is : %s\n", name)
	fmt.Printf("my hobbies is : %s\n", hobbiesAsString)
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}

// alias
type FilterCallback func(string) bool

func filter(data []string, callback FilterCallback) []string {
	var result []string
	for _, each := range data {
		var filtered = callback(each)
		if filtered {
			result = append(result, each)
		}
	}
	return result
}
