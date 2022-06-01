package main

import (
	"encoding/base64"
	"fmt"
	"log"

	_ "embed"
	// "encoding/json"
)

//go:embed icon.png
var input []byte

func main() {
	// input := []byte("hello world")
	fmt.Println("png size:", len(input))
	fmt.Println("///////////////////////////////input")
	fmt.Println(input)

	fmt.Println("///////////////////////////////演示base64编码")
	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Println(encodeString)

	fmt.Println("///////////////////////////////对上面的编码结果进行base64解码")
	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(decodeBytes))

	// // 如果要用在url中，需要使用URLEncoding
	// uEnc := base64.URLEncoding.EncodeToString([]byte(input))
	// fmt.Println(uEnc)

	// uDec, err := base64.URLEncoding.DecodeString(uEnc)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(string(uDec))
}
