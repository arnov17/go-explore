package main

import "fmt"

func main(){
	fmt.Println("Hello World")
	fmt.Println("Satu =", 1)
	fmt.Println("Dua =", 2)
	fmt.Println("Tiga Koma Lima =", 3.5)
	fmt.Println("benar =", true)
	fmt.Println("benar =", false)
	fmt.Println(len("Arnov"))
	fmt.Println("Arnov"[0]) // out punya byte, jika menjadi string prlu convert
}

//compile kdengan 'go build <filename>.go
// lihat hasil compile ls -l, then ./<filename>
// ketika app sudah jadi perlu di compile

// go run <filename>.go