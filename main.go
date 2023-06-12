package main

import (
	"fmt"
	"os"
)

func main() {
	text := os.Args[1]
	fileName := os.Args[2]

	fmt.Println(text, fileName)

	search(text, fileName)

}
