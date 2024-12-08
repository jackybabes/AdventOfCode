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

type Calc struct {
	total  int
	nums   []int
	totals []int
}

func readInputIntoArrays(filename string) ([]Calc, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}
	defer file.Close()

	var calcs []Calc
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		line := scanner.Text()
		lineSplit := strings.Split(line, ":")
		var calc Calc
		total, _ := strconv.Atoi(lineSplit[0])
		calc.total = total
		numsStrings := strings.Split(strings.Trim(lineSplit[1], " "), " ")
		for _, num := range numsStrings {
			nint, _ := strconv.Atoi(num)
			calc.nums = append(calc.nums, nint)
		}
		calcs = append(calcs, calc)
	}

	return calcs, nil
}

func main() {
	filename := "input.txt"

	calcs, err := readInputIntoArrays(filename)
	if err != nil {
		log.Fatalf("file error: %v", err)
	}

	// operators := [4]string{"+", "-", "*", "/"}

	var validSum int
	for _, calc := range calcs {
		log.Println(calc)

		calc.totals = append(calc.totals, calc.nums[0])
		calc.nums = calc.nums[1:]

		for _, num := range calc.nums {
			initalLenOfTotals := len(calc.totals)
			for _, intTotal := range calc.totals {

				add := intTotal + num
				// if add <= calc.total {
				calc.totals = append(calc.totals, add)
				// }

				mult := intTotal * num
				// if mult <= calc.total {
				calc.totals = append(calc.totals, mult)
				// }

				concat := strconv.FormatInt(int64(intTotal), 10) + strconv.FormatInt(int64(num), 10)
				concatInt, _ := strconv.Atoi(concat)
				calc.totals = append(calc.totals, concatInt)
			}
			// remove earlier totals
			calc.totals = calc.totals[initalLenOfTotals:]
		}

		if slices.Contains(calc.totals, calc.total) {
			log.Println("Valid")
			validSum += calc.total
		}

		// break
	}
	log.Printf("Valid Sum = %v", validSum)

}

// 2 -> 3
// 3 -> 9
// 4 -> 27
