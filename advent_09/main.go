package main

import (
	"log"
	"slices"
	"strconv"
)

func sToI(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func unpackFiles(dm string) []string {
	var unpacked []string
	id := 0
	freeSpace := false
	for _, data := range dm {
		dataString := string(data)
		if !freeSpace {
			for range sToI(dataString) {
				unpacked = append(unpacked, strconv.Itoa(id))
			}
			id++
			freeSpace = !freeSpace
		} else {
			for range sToI(dataString) {
				unpacked = append(unpacked, ".")
			}
			freeSpace = !freeSpace
		}
	}
	return unpacked
}

func checkFreeSpace(s []string) bool {
	return slices.Contains(s, ".")
}

func moveFinalDataPointIntoFirstFreeSpace(s []string) []string {
	finalData := s[len(s)-1]
	s = s[:len(s)-1]
	for i, r := range s {
		if r == "." {
			s[i] = finalData
			break
		}
	}
	return s
}

func calcCheckSum(s []string) int {
	var sum int
	for i, r := range s {
		digit := sToI(string(r))
		sum += i * digit
	}
	return sum
}

func main() {
	filename := "input.txt"
	diskMapPacked := readInputIntoStr(filename)

	log.Printf("Len of diskmap packed: %v", len(diskMapPacked))

	diskMap := unpackFiles(diskMapPacked)

	log.Printf("Len of diskmap unpacked: %v", len(diskMap))

	for checkFreeSpace(diskMap) {
		diskMap = moveFinalDataPointIntoFirstFreeSpace(diskMap)
		// log.Println(diskMap)
	}

	log.Printf("Len of diskmap sorted: %v", len(diskMap))

	log.Println(calcCheckSum(diskMap))

}
