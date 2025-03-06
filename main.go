package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetFileContent(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err.Error())
		return ""
	}
	return string(content)
}

func GetFactors(n uint64) [2]uint64 {
	var factors [2]uint64
	var i uint64

	for i = 2; i < n; i++ {
		if n%i == 0 {
			factors[0] = i
			factors[1] = n / i
			break
		}
	}
	return factors
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./factors <filename>")
		os.Exit(1)
	}
	file := os.Args[1]
	fileContent := GetFileContent(file)
	if fileContent == "" {
		fmt.Fprintln(os.Stderr, "Error: file is empty")
		os.Exit(1)
	}
	contents := strings.Split(fileContent, "\n")
	for _, content := range contents {
		if content == "" {
			continue
		}
		num, _ := strconv.Atoi(content)
		factors := GetFactors(uint64(num))
		fmt.Printf("%v=%v*%v\n", content, factors[1], factors[0])
	}
}
