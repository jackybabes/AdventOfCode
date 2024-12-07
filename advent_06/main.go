package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"sync"
	"time"
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
	var matrix [][]string
	for _, line := range input {
		matrix = append(matrix, strings.Split(line, ""))
	}

	return matrix, nil

}

func printMatrix(ws [][]string) {
	for _, line := range ws {
		log.Println(line)
	}
}

type xyMap struct {
	matrix [][]string
}

func (m xyMap) get(x, y int) string {
	line := m.matrix[len(m.matrix)-1-y]
	point := line[x]
	return point
}
func (m xyMap) set(x, y int, s string) {
	line := m.matrix[len(m.matrix)-1-y]
	line[x] = s
}
func (m xyMap) print() {
	printMatrix(m.matrix)
}

func (m *xyMap) findFirst(icons []string) (int, int) {
	for x := range m.matrix[0] {
		for y := range m.matrix {
			if slices.Contains(icons, m.get(x, y)) {
				return x, y
			}
		}
	}
	return -1, -1
}

func (m *xyMap) checkInBounds(x, y int) bool {
	if x >= 0 && x < len(m.matrix[0]) && y >= 0 && y < len(m.matrix) {
		return true
	}
	return false
}

func (m *xyMap) countPositions() int {
	var count int
	for x := range m.matrix[0] {
		for y := range m.matrix {
			if m.get(x, y) == "X" {
				count++
			}
		}
	}
	return count
}

type GuardLog struct {
	x   int
	y   int
	dir string
}

func runSimulation(xy *xyMap, guardIcons []string, guardDirections map[string][2]int) bool {
	// find gaurd and store location and direction
	guardX, guardY := xy.findFirst(guardIcons)
	guardDirIcon := xy.get(guardX, guardY)

	var guardLogs []GuardLog

	// loop:
	for {
		// log.Printf("Guard at: (%v, %v), travelling %v", guardX, guardY, guardDirIcon)
		// make guard log
		guardLog := GuardLog{guardX, guardY, guardDirIcon}

		// check in guard log contained in history
		if slices.Contains(guardLogs, guardLog) {
			log.Println("Infinate loop")
			return true
		}
		// store guard log
		guardLogs = append(guardLogs, guardLog)

		// mark current square with X
		xy.set(guardX, guardY, "X")

	testDirectionAfterTurn:
		// check direction travelling (^>v<)
		guardDir := guardDirections[guardDirIcon]

		// check if possible to take step in direction
		nextTileX, nextTileY := guardX+guardDir[0], guardY+guardDir[1]

		if !xy.checkInBounds(nextTileX, nextTileY) {
			// if walk off map end loop
			// log.Println("Heading off map")
			// xy.print()
			break
		}

		// if obstical (#), rotate direction 90 degrees clockwise
		nextTileIcon := xy.get(nextTileX, nextTileY)
		// log.Println(nextTileIcon)
		if nextTileIcon == "#" {
			// log.Println("Hit obstical, turning...")
			indexOfDirection := slices.Index(guardIcons, guardDirIcon)
			guardDirIcon = guardIcons[(indexOfDirection+1)%len(guardIcons)]
			goto testDirectionAfterTurn
		}

		// take step
		guardX = nextTileX
		guardY = nextTileY

		// write gaurd icon to map/ update guard xy
		xy.set(guardX, guardY, guardDirIcon)
		// repeat and print
		// xy.print()
		// time.Sleep(time.Millisecond * 200)
	}

	// count X on map
	// log.Println(xy.countPositions())
	return false
}

func deepCopyMatrix(matrix [][]string) [][]string { // GPT
	// Create a new outer slice with the same length as the input matrix
	newMatrix := make([][]string, len(matrix))
	for i := range matrix {
		// Create a new inner slice for each row and copy the contents
		newMatrix[i] = make([]string, len(matrix[i]))
		copy(newMatrix[i], matrix[i])
	}
	return newMatrix
}

func createMapsWithExtraObstical(matrix [][]string, guardIcons []string) []xyMap {

	xyInitial := xyMap{slices.Clone(matrix)}
	guardX, guardY := xyInitial.findFirst(guardIcons)

	var xyMaps []xyMap
	for x := range matrix[0] {
		for y := range matrix {
			xy := xyMap{deepCopyMatrix(matrix)}
			if x == guardX && y == guardY {
				continue
			}
			xy.set(x, y, "#")
			xyMaps = append(xyMaps, xy)
		}
	}
	return xyMaps
}

func main() {
	filename := "input.txt"
	guardIcons := []string{"^", ">", "v", "<"}
	guardDirections := map[string][2]int{
		"^": {0, 1},  // Up
		">": {1, 0},  // Right
		"v": {0, -1}, // Down
		"<": {-1, 0}, // Left
	}

	guardMap, err := readInputIntoMatrix(filename)
	if err != nil {
		log.Fatalf("Input Error: %v", err)
	}
	log.Println("Making Maps")

	start := time.Now()
	arrayOfMaps := createMapsWithExtraObstical(guardMap, guardIcons)
	log.Printf("Took %v to create %v maps", time.Since(start), len(arrayOfMaps))

	var infCount int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, testMap := range arrayOfMaps {
		// log.Printf("Testing %v/%v Map", i, len(arrayOfMaps))
		wg.Add(1)
		go func(m xyMap) {
			defer wg.Done()
			if runSimulation(&m, guardIcons, guardDirections) {
				mu.Lock()
				infCount++
				mu.Unlock()
			}
		}(testMap)

		// log.Println(testMap.countPositions())
	}
	wg.Wait()
	log.Println(infCount)

}
