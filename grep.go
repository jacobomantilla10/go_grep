package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"regexp"
)

func searchRecursively(s string, d string) {
	var dir string

	if d == "." {
		dir, _ = os.Getwd()
	} else {
		dir = d
	}

	fileSystem := os.DirFS(dir)
	fmt.Println(fileSystem)

	output := ""

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
		}

		if d.Name() == "grep" {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		output += searchFile(s, d.Name())

		if output != "" {
			output += "\n"
		}

		return nil
	})

	fmt.Println(output)
}

func searchFile(s string, f string) string {
	file, err := os.Open(f)

	if err != nil {
		return ""
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

	file.Close()

	return output
}
