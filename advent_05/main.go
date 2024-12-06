package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// ConvertStringToIntSlice converts a comma-separated string to a slice of integers. GPT
func ConvertStringToIntSlice(input string) []int {
	// Split the string by commas
	parts := strings.Split(input, ",")
	result := make([]int, len(parts))

	// Convert each part to an integer
	for i, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part)) // Ensure no leading/trailing spaces
		result[i] = num
	}
	return result
}

func readInputIntoArrays(filename string) ([]string, []string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, nil, err
	}
	defer file.Close()
	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, nil, err
	}

	indexOfGap := slices.Index(input, "")

	pageOrderingRulesStr := input[:indexOfGap]
	pagesToProduceStr := input[indexOfGap+1:]

	return pageOrderingRulesStr, pagesToProduceStr, nil
}

func main() {
	// log.Println("Hello World")

	filename := "input.txt"
	pageOrderingRules, pagesToProduce, err := readInputIntoArrays(filename)
	if err != nil {
		log.Fatalf("Input Error: %v", err)
	}

	// log.Println(pageOrderingRules)
	// log.Println(pagesToProduce)

	keyMustBeBeforeMap := make(map[int][]int)
	keyMustBeAfterMap := make(map[int][]int)
	for _, rule := range pageOrderingRules {
		beforeAndAfter := strings.Split(rule, "|")
		before, _ := strconv.Atoi(beforeAndAfter[0])
		after, _ := strconv.Atoi(beforeAndAfter[1])
		keyMustBeBeforeMap[before] = append(keyMustBeBeforeMap[before], after)
		keyMustBeAfterMap[after] = append(keyMustBeAfterMap[after], before)
	}

	var manuals [][]int
	for _, manual := range pagesToProduce {
		manuals = append(manuals, ConvertStringToIntSlice(manual))
	}

	var goodManuals [][]int

ManualsLoop:
	for _, manual := range manuals {
		log.Println(manual)
		for i, page := range manual {
			pagesBefore := manual[:i]

			for _, pageBefore := range pagesBefore {
				if !slices.Contains(keyMustBeAfterMap[page], pageBefore) {
					log.Println("Bad times")
					continue ManualsLoop
				}
			}

			pagesAfter := manual[i+1:]

			for _, pageAfter := range pagesAfter {
				if !slices.Contains(keyMustBeBeforeMap[page], pageAfter) {
					log.Println("Bad times")
					continue ManualsLoop
				}
			}

			_, _, _, _ = i, pagesAfter, pagesBefore, page

		}
		log.Println("Success")
		goodManuals = append(goodManuals, manual)
	}

	log.Println(goodManuals)

	var total int

	for _, gm := range goodManuals {
		indexOfHalf := (len(gm) / 2)

		total += gm[indexOfHalf]
	}

	log.Println(total)

}
