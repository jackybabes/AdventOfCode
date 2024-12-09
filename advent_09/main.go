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

func findFreeSpaceLen(s []string, len int) int {
	var freeSpaces int
	for i, e := range s {
		if e == "." {
			freeSpaces++
			if freeSpaces == len {
				return i - len + 1
			}
			continue
		}
		freeSpaces = 0
	}
	return -1
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

	// Part 2 (THIS IS A MESS... but it works)
	var fileLen int
	var id string
	for i := range diskMap {
		i = len(diskMap) - i - 1
		if id == "" {
			id = diskMap[i]
			fileLen++
			continue
		}
		if id == diskMap[i] {
			fileLen++
			continue
		}

		indexOfFreeSpace := findFreeSpaceLen(diskMap, fileLen)
		// log.Println(indexOfFreeSpace)
		if indexOfFreeSpace > 0 && indexOfFreeSpace < i {
			for j := range fileLen {
				diskMap[indexOfFreeSpace+j] = diskMap[i+j+1]
				diskMap[i+j+1] = "."
			}
		}

		fileLen = 1
		id = diskMap[i]

		continue

	}

	// Part One
	// for checkFreeSpace(diskMap) {
	// 	diskMap = moveFinalDataPointIntoFirstFreeSpace(diskMap)
	// 	// log.Println(diskMap)
	// }

	// log.Println(diskMap)

	log.Printf("Len of diskmap sorted: %v", len(diskMap))
	log.Printf("Checksum of DiskMap: %v", calcCheckSum(diskMap))

}
