package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var str = "124"
	// var num, err = strconv.Atoi(str)

	// if err == nil {
	// 	fmt.Println(num) // 124
	// }

	// var num1 = 124
	// var str1 = strconv.Itoa(num1)

	// fmt.Println(str1) // "124"

	// var str = "124"
	// var num, err = strconv.ParseInt(str, 10, 64)

	// if err == nil {
	// 	fmt.Println(num) // 124
	// }

	// var str = "1010"
	// var num, err = strconv.ParseInt(str, 2, 8)

	// if err == nil {
	// 	fmt.Println(num) // 10
	// }

	// var str = "24.12"
	// var num, err = strconv.ParseFloat(str, 32)

	// if err == nil {
	// 	fmt.Println(num) // 24.1200008392334
	// }

	// var str = "true"
	// var bul, err = strconv.ParseBool(str)

	// if err == nil {
	// 	fmt.Println(bul) // true
	// }

	var bul = true
	var str = strconv.FormatBool(bul)

	fmt.Println(str) // true
}
