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

func readInputIntoArrays(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}
	defer file.Close()

	// var list1, list2 []int
	var lists [][]int

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Print each line
		// fmt.Println(scanner.Text())
		stringNumsInArray := strings.Split(scanner.Text(), " ")

		var intNumsInArray []int
		for _, str := range stringNumsInArray {
			i, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			intNumsInArray = append(intNumsInArray, i)
		}

		lists = append(lists, intNumsInArray)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, err
	}

	return lists, nil
}

func checkIncOrDec(report []int) bool {
	sorted_report := slices.Clone(report)
	slices.Sort(sorted_report)
	reverse_sorted_report := slices.Clone(sorted_report)
	slices.Reverse(reverse_sorted_report)
	return slices.Compare(report, sorted_report) == 0 || slices.Compare(report, reverse_sorted_report) == 0
}

func checkAdjacentLevelsBetweenOneAndThree(report []int) bool {
	var differnce []int
	for i := range len(report) - 1 {
		d := abs(report[i] - report[i+1])
		differnce = append(differnce, d)
	}
	for _, n := range differnce {
		if n < 1 || n > 3 {
			return false
		}
	}
	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {

	reports, err := readInputIntoArrays("input.txt")
	if err != nil {
		log.Fatalf("Input Error: %v", err)
	}

	var trueReports int

	for _, report := range reports {
		fmt.Println(report)

		// fmt.Println(checkIncOrDec(report))

		// fmt.Println(checkAdjacentLevelsBetweenOneAndThree(report))

		fmt.Println(checkIncOrDec(report) && checkAdjacentLevelsBetweenOneAndThree(report))

		if checkIncOrDec(report) && checkAdjacentLevelsBetweenOneAndThree(report) {
			trueReports++
		}
	}

	fmt.Printf("True Reports: %v", trueReports)
}
