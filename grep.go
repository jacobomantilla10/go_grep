package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func search(s string, fName string) {
	output := searchFile(s, fName)
	fmt.Println(output)
}

func searchFile(s string, f string) string {
	file, err := os.Open(f)

	if err != nil {
		panic(1)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	output := ""
	replacement := "\x1b[31m" + s + "\x1b[0m"
	re, _ := regexp.Compile(s)

	for scanner.Scan() {
		if re.MatchString(scanner.Text()) {
			output += re.ReplaceAllString(scanner.Text(), replacement) + "\n"
		}
	}

	//outputBytes := re.ReplaceAllString(string(f), replacement)

	return output
}
