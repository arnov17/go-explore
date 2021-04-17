package main

import "fmt"

func main() {
	type noKtp string
	type married bool

	var noKtpEko noKtp = '12342424234' // menjadi type string
	var marriedStatus married = true // menjadi type bool
}