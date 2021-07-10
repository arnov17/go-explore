package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println(chicken)                       // map[februari:40 januari:50]
	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	// var data map[string]int
	// data["one"] = 1
	// akan muncul error! karena default vue map adalah "nil", perlu inisialisasi

	// var cat map[int]string
	cat := map[int]string{}
	cat[11] = "rokes"
	cat[22] = "takil"
	fmt.Println(cat)     // map[11:rokes 22:takil]
	fmt.Println(cat[11]) // rokes
	/*
		// cara horizontal
		var chicken1 = map[string]int{"januari": 50, "februari": 40}

		// cara vertical
		var chicken2 = map[string]int{
			"januari":  50,
			"februari": 40,
		}
	*/

	// var chickens = map[string]int{
	// 	"januari":  50,
	// 	"februari": 40,
	// 	"maret":    34,
	// 	"april":    67,
	// }
	// fmt.Println(len(chickens)) // 1
	// fmt.Println(chickens)

	// for key, val := range chickens {
	// 	fmt.Println(key, "  \t:", val)
	// }

	// delete(chickens, "maret")
	// fmt.Println(len(chickens)) // 1
	// fmt.Println(chickens)

	var months = map[string]int{
		"januari":  90,
		"februari": 88,
		"mei":      77,
		"juni":     80,
	}
	var value, isExist = months["mei"]
	fmt.Println(isExist)
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not found")
	}

	// var chickens = []map[string]string{
	// 	map[string]string{"name": "blue", "gender": "male"},
	// 	map[string]string{"name": "echo", "gender": "female"},
	// 	map[string]string{"name": "delta", "gender": "male"},
	// }

	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	var data = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}

	for _, d := range data {
		fmt.Println(d)
	}

}
