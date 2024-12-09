package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Coord struct {
	x int
	y int
}

// type Antenna struct {
// 	pos  Coord
// 	freq string
// }

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
	var matrix [][]string
	for _, line := range input {
		matrix = append(matrix, strings.Split(line, ""))
	}

	return matrix, nil

}

func printMatrix(m [][]string) {
	for _, line := range m {
		log.Println(line)
	}
}

func makePairs(ants []Coord) [][2]Coord {
	var pairs [][2]Coord

	for _, a := range ants {
		for _, b := range ants {
			if a != b {
				pair := [2]Coord{a, b}
				pairRev := [2]Coord{a, b}
				pairSwitched := [2]Coord{b, a}
				pairRevSwitched := [2]Coord{b, a}
				if !slices.Contains(pairs, pair) &&
					!slices.Contains(pairs, pairRev) &&
					!slices.Contains(pairs, pairSwitched) &&
					!slices.Contains(pairs, pairRevSwitched) {
					pairs = append(pairs, pair)
				}
			}
		}
	}
	return pairs
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func makeAntiNodes(pair [2]Coord) [2]Coord {
	pointA := pair[0]
	pointB := pair[1]

	xDistance := abs(pointA.x - pointB.x)
	yDistance := abs(pointA.y - pointB.y)

	antiA := Coord{}
	antiB := Coord{}

	if pointA.x < pointB.x {
		antiA.x = pointA.x - xDistance
		antiB.x = pointB.x + xDistance
	} else {
		antiA.x = pointA.x + xDistance
		antiB.x = pointB.x - xDistance
	}

	if pointA.y < pointB.y {
		antiA.y = pointA.y - yDistance
		antiB.y = pointB.y + yDistance
	} else {
		antiA.y = pointA.y + yDistance
		antiB.y = pointB.y - yDistance
	}

	// antiA := Coord{pointA.x + xDistance, pointA.y + yDistance}
	// antiB := Coord{pointB.x - xDistance, pointB.y - yDistance}

	return [2]Coord{antiA, antiB}

}

func main() {
	filename := "input.txt"

	matrix, err := readInputIntoMatrix(filename)
	if err != nil {
		log.Fatalf("file error: %v", err)
	}

	antennaMap := xyMap{matrix}
	antennaMap.print()

	antennas := antennaMap.collectAntenna()

	// log.Println(ants)
	for freq, coords := range antennas {
		log.Printf("Antennas of freq '%v' at postions %v", freq, coords)
		pairs := makePairs(coords)
		log.Println(pairs)

		for _, pair := range pairs {
			antiNodes := makeAntiNodes(pair)
			log.Println(antiNodes)
			for _, node := range antiNodes {
				if antennaMap.checkInBounds(node) {
					// if antennaMap.get(node) == "." {
					antennaMap.set(node, "#")
					// }
				}
			}

		}

	}

	antennaMap.print()
	log.Println(antennaMap.countPositions("#"))

}
