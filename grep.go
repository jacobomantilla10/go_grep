package main

import (
	"fmt"
	"os"
)

func search(s string, fName string) {
	fmt.Println(s, fName)

	file, err := os.ReadFile(fName)

	if err != nil {
		panic(err)
	}

	//fmt.Println(string(file))
	//fmt.Println("file bytes: ", file, "string bytes: ", []byte(s))
	num := getNumOccurrences([]byte(s), file)
	fmt.Println(num)
}

func getNumOccurrences(s []byte, f []byte) int {
	numOccurrences := 0
	j := 0

	for _, val := range f {
		if j >= len(s) {
			numOccurrences++
			j = 0
		}

		if val == s[j] {
			j++
		}
	}
	return numOccurrences
}
