package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(10)
	fmt.Println("random ke-1:", rand.Int()) // 5221277731205826435
	fmt.Println("random ke-2:", rand.Int()) // 3852159813000522384
	fmt.Println("random ke-3:", rand.Int()) // 8532807521486154107

	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())

	fmt.Println(randomString(10))

}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	// fmt.Printf("%T\n", b) // []int32 rune
	return string(b)
}
