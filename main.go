package main

import (
	"os"
)

func main() {
	text := os.Args[1]
	fileName := os.Args[2]

	search(text, fileName)

}
