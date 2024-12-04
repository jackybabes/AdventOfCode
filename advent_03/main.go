package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

func readInputIntoString(filename string) string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}
	return string(bytes)
}

func subMatchesMultply(subMatches [][]string) int {
	var total int
	for _, subMatch := range subMatches {
		// log.Println(subMatch)
		i, err := strconv.Atoi(subMatch[1])
		if err != nil {
			log.Fatalf("Error converting str to int: %v\n", err)
		}
		j, err := strconv.Atoi(subMatch[2])
		if err != nil {
			log.Fatalf("Error converting str to int: %v\n", err)
		}
		// log.Println(i * j)
		total += i * j
	}
	return total
}

func main() {
	filename := "input.txt"
	str := readInputIntoString(filename)

	log.Println(str)

	matchMuls, err := regexp.Compile(`(?:mul\()(\d{1,3})(?:,)(\d{1,3})(?:\))`)
	if err != nil {
		log.Fatalf("Error compiling regex: %v\n", err)
	}

	matchDoDont, err := regexp.Compile(`don't\(\)(.|\n|\r|\t)*?do\(\)`)
	if err != nil {
		log.Fatalf("Error compiling regex: %v\n", err)
	}
	for {
		indexOfDont := matchDoDont.FindStringIndex(str)
		if indexOfDont == nil {
			break
		}
		str = str[:indexOfDont[0]] + str[indexOfDont[1]:]
	}

	matchFinalDont, err := regexp.Compile(`don't\(\)(.|\n)*`)
	if err != nil {
		log.Fatalf("Error compiling regex: %v\n", err)
	}
	finalDontIndex := matchFinalDont.FindStringIndex(str)
	if finalDontIndex != nil {
		str = str[:finalDontIndex[0]]
	}

	log.Println(str)

	subMatches := matchMuls.FindAllStringSubmatch(str, -1)
	// log.Println(subMatches)

	addedMatches := subMatchesMultply(subMatches)
	log.Printf("Total: %v", addedMatches)

}
