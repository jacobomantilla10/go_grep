package main

import (
	"fmt"
	"os"
	"regexp"
)

func search(s string, fName string) {
	file, err := os.ReadFile(fName)

	if err != nil {
		panic(err)
	}

	output := searchFile(s, file)
	fmt.Println(output)
}

func searchFile(s string, f []byte) string {
	re, _ := regexp.Compile(s)

	replacement := "\x1b[31m" + s + "\x1b[0m"
	outputBytes := re.ReplaceAllString(string(f), replacement)

	return outputBytes
}
