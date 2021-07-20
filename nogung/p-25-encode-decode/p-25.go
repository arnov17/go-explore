package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// var data = "john wick"

	// var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	// fmt.Println("encoded:", encodedString)

	// var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	// fmt.Println("decoded:", decodedByte)
	// var decodedString = string(decodedByte)
	// fmt.Println("decoded:", decodedString)

	var data = "john wick"

	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodedString = string(encoded)
	fmt.Println(encodedString)

	var decocded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decocded, []byte(decocded))
	if err != nil {
		fmt.Println(err.Error())
	}
	var decodedString = string(decocded)
	fmt.Println(decodedString)

	var dataURL = "https://kalipare.com/"

	var encodedStringURL = base64.URLEncoding.EncodeToString([]byte(dataURL))
	fmt.Println(encodedStringURL)

	var decodedByte, _ = base64.URLEncoding.DecodeString(encodedStringURL)
	var decodedStringURL = string(decodedByte)
	fmt.Println(decodedStringURL)
}
