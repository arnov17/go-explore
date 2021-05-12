package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.7)) // 2
	fmt.Println(math.Round(1.3)) // 1
	fmt.Println(math.Floor(1.7)) // 1
	fmt.Println(math.Ceil(1.3))  // 2

	fmt.Println(math.Max(10, 25))  // 25
	fmt.Println(math.Min(1.3, 34)) // 1.3

}
