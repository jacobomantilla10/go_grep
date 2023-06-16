package main

import (
	"flag"
	"fmt"
)

var rFlag = flag.Bool(
	"r",
	false,
	"Use -r in order to search in current directory and all subdirectories",
)

func main() {
	flag.Parse()

	args := flag.Args()
	pattern := args[0]
	//fileName := args[1]

	if *rFlag {
		directory := args[1]
		searchRecursively(pattern, directory)
	} else {
		fileName := args[1]
		fmt.Println(searchFile(pattern, fileName))
	}
}
