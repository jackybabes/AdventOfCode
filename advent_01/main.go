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

func readInputIntoArrays(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, nil, err
	}
	defer file.Close()

	// var list1, list2 []int
	lists := make([][]int, 2)

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Print each line
		// fmt.Println(scanner.Text())
		stringNumsInArray := strings.Split(scanner.Text(), "   ")

		for i, str := range stringNumsInArray {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return nil, nil, err
			}
			lists[i] = append(lists[i], num)
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, nil, err
	}

	return lists[0], lists[1], nil
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {

	filename := "input.txt"

	list1, list2, err := readInputIntoArrays(filename)
	if err != nil {
		log.Fatalf("file error: %v", err)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	var difference []int
	for i := range list1 {
		difference = append(difference, abs(list1[i]-list2[i]))
	}

	fmt.Println(list1)
	fmt.Println(list2)
	fmt.Println(difference)

	var distance int
	for _, n := range difference {
		distance += n
	}

	fmt.Println(distance)

	similarity := 0
	for _, num1 := range list1 {

		// how often does n appear in list2
		count := 0
		for _, num2 := range list2 {
			if num1 == num2 {
				count++
			}
		}
		similarity += count * num1
	}

	fmt.Println(similarity)

}
