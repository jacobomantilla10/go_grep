package main

import (
	"fmt"
	"os"
	"strings"
)

func search(s string, fName string) {
	file, err := os.ReadFile(fName)

	if err != nil {
		panic(err)
	}

	output := formatOutput(s, strings.Split(string(file), " "))
	fmt.Println(output)
}

func formatOutput(s string, f []string) string {
	output := []string{}

	for _, val := range f {
		if val == s {
			output = append(output, "\033[31m", val, "\033[0m")
		} else {
			output = append(output, val)
		}
	}
	return strings.Join(output, " ")
}
