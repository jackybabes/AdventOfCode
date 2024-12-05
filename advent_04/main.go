package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInputIntoMatrix(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}
	defer file.Close()
	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, err
	}
	// return input, nil

	var matrix [][]string
	for _, line := range input {
		matrix = append(matrix, strings.Split(line, ""))
	}

	return matrix, nil

}

func printWS(ws [][]string) {
	for _, line := range ws {
		log.Println(line)
	}
}

func checkStringExistsInDirection(wordSearch [][]string, i, j int, line []string, wordToFind string, v [2]int) bool {
	verticalMax := i + ((len(wordToFind) - 1) * v[0])
	horizontalMax := j + ((len(wordToFind) - 1) * v[1])
	if 0 <= verticalMax && verticalMax < len(wordSearch) && 0 <= horizontalMax && horizontalMax < len(line) {
		return true
	}
	return false
}

func makeEmpty() [][]string {
	rows, cols := 10, 10
	matrix := make([][]string, rows)
	for i := range matrix {
		matrix[i] = make([]string, cols)
		for j := range matrix[i] {
			matrix[i][j] = "."
		}
	}
	return matrix
}

func main() {
	filename := "input.txt"
	wordToFind := "MAS"
	directions := map[string][2]int{
		// "e":  {0, 1},
		"se": {1, 1},
		// "s":  {1, 0},
		"sw": {1, -1},
		// "w":  {0, -1},
		"nw": {-1, -1},
		// "n":  {-1, 0},
		"ne": {-1, 1},
	}
	// emptyMatrix := makeEmpty()

	wordSearch, err := readInputIntoMatrix(filename)
	if err != nil {
		log.Fatalf("Input Error: %v", err)
	}
	printWS(wordSearch)
	// printWS(emptyMatrix)

	var wordFound [][2]int
	var timesWordFound int

	for i, line := range wordSearch {
		for j, char := range line {
			if char == string(wordToFind[0]) {
				log.Printf("Found X at (%v, %v)", i, j)

				for k, v := range directions {
					_ = k
					// log.Printf("Checking we can get a string of len: %v in direction: %v", len(wordToFind), k)
					if !checkStringExistsInDirection(wordSearch, i, j, line, wordToFind, v) {
						// log.Println("String does not Exist within wordsearch")
						continue
					}

					var stringInDirection string
					for x := range len(wordToFind) {
						stringInDirection += wordSearch[i+(x*v[0])][j+(x*v[1])]
					}

					// log.Println(stringInDirection)

					if stringInDirection == wordToFind {
						// emptyMatrix[i][j] = "X"
						// emptyMatrix[i+(1*v[0])][j+(1*v[1])] = "M"
						// emptyMatrix[i+(2*v[0])][j+(2*v[1])] = "A"
						// emptyMatrix[i+(3*v[0])][j+(3*v[1])] = "S"
						log.Println("Success")
						wordFound = append(wordFound, [2]int{i + (1 * v[0]), j + (1 * v[1])})
						// timesWordFound++
					}
				}

			}
		}
	}

	numberOfCrossingWord := make(map[[2]int]int)

	for _, word := range wordFound {
		numberOfCrossingWord[word]++
	}

	for _, v := range numberOfCrossingWord {
		if v > 1 {
			timesWordFound++
		}
	}

	log.Println(timesWordFound)

}
