package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := []byte("Hello, Base64 Encoding")

	// Encode Base64
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	// Decode from Base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Decoding error:", err)
	}
	fmt.Println("Decode:", string(decoded))

	// URL safe, avoid '/' and '+'
	data2 := []byte("He~o, Base64 Encoding")

	urlSafeEncoded := base64.StdEncoding.EncodeToString(data2)
	fmt.Println("URL Safe Encoded: ", urlSafeEncoded)

	// plainStr := "Hello World"

	// encodedStr := base64.StdEncoding.EncodeToString([]byte(plainStr))
	// fmt.Println(encodedStr)

	// decodedStrAsByteSlice, err := base64.StdEncoding.DecodeString(encodedStr)
	// if err != nil {
	// 	fmt.Println("Decoding error:", err)
	// 	return
	// }
	// fmt.Println("Decode:", string(decodedStrAsByteSlice))

}
