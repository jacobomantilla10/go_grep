package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"regexp"
)

func search(s string, fName string) {
	//output := searchFile(s, fName, false)
	//fmt.Println(output)
	findFile(fName)
}

func findFile(f string) {
	// Check every file in the current directory
	// Iterate through all subdirectories, while checking every file in each subdirectory
	// Do the above step recursively
	wd, err := os.Getwd()

	if err != nil {
		panic(1)
	}

	fileSystem := os.DirFS(wd)
	fmt.Println(fileSystem)

	output := ""
	// Need to fix the way this is called, this works right now for testing purposes
	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
		}

		if d.Name() == "grep" {
			return nil
		}

		output += searchFile("Linux", d.Name(), d.IsDir())

		if output != "" {
			output += "\n"
		}
		//fmt.Println("Linux", d.Name(), d.IsDir())

		return nil
	})

	fmt.Println(output)
}

func searchFile(s string, f string, isDir bool) string {
	if isDir {
		return ""
	}

	file, err := os.Open(f)

	if err != nil {
		//panic(1)
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
