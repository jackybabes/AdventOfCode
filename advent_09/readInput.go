package main

import (
	"bufio"
	"os"
)

func readInputIntoStr(filename string) string {
	file, _ := os.Open(filename)
	defer file.Close()
	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input[0]
}
