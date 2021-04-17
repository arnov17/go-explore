package main

import "fmt"

func main() {
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32) // akan  minus, karena nilainya diatas range type int8
 
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)


	var name string= "arnov"
	var e byte= name[0]
	var eString string = string(e)
	
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}