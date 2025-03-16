package main

import (
	"codewars-go-solutions/solutions"
	"fmt"
)

func main() {
	encode := solutions.Encode([]byte("test"))
	fmt.Println(encode)
}
