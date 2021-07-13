package main

import (
	"fmt"
	"math"
	"strings"
)

// type hitung interface {
// 	luas() float64
// 	keliling() float64
// }

// type lingkaran struct {
// 	diameter float64
// }

// type persegi struct {
// 	sisi float64
// }

// func (l lingkaran) jariJari() float64 {
// 	return l.diameter / 2
// }

// func (l lingkaran) luas() float64 {
// 	return math.Pi * math.Pow(l.jariJari(), 2)
// }

// func (l lingkaran) keliling() float64 {
// 	return math.Pi * l.diameter
// }

// func (p persegi) luas() float64 {
// 	return math.Pow(p.sisi, 2)
// }

// func (p persegi) keliling() float64 {
// 	return p.sisi * 4
// }

// interface embed
type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func main() {

	var bangunRuang hitung = &kubus{4}

	fmt.Println("===== kubus")
	fmt.Println("luas      :", bangunRuang.luas())
	fmt.Println("keliling  :", bangunRuang.keliling())
	fmt.Println("volume    :", bangunRuang.volume())

	// var secret interface{}

	// secret = "ethan"
	// fmt.Println(secret)

	// secret = []string{"apple", "manggo", "banana"}
	// fmt.Println(secret)

	// secret = 12.8

	// fmt.Println(secret)

	// var data map[string]interface{}

	// data = map[string]interface{}{
	// 	"name":      "ethan hunt",
	// 	"grade":     2,
	// 	"breakfast": []string{"apple", "manggo", "banana"},
	// }
	// fmt.Println(data)

	var secret interface{}

	secret = 2
	var number = secret.(int) * 10

	fmt.Println(secret, "multiplied by 10 is", number)

	secret = []string{"apple", "manggo", "banana"}
	var fruits = strings.Join(secret.([]string), ", ")

	fmt.Println(fruits, "is my favorite fruits")

	type person struct {
		name string
		age  int
	}

	var secret2 interface{} = &person{name: "wick", age: 18}
	var name = secret2.(*person).name
	fmt.Println(secret2)
	fmt.Println(name)

	var person2 = []map[string]interface{}{
		{"name": "wick", "age": 29},
		{"name": "john", "age": 31},
		{"name": "doe", "age": 22},
		{"name": "mina", "age": 26},
	}

	for _, each := range person2 {
		fmt.Println(each["name"], "age is", each["age"])
	}

	var fruits2 = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}

	for _, each := range fruits2 {
		fmt.Println(each)
	}

}
