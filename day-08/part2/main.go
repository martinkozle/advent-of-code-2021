package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func readLines(file string) []string {
	data, err := os.ReadFile(file)
	check(err)
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func containsAllRunes(s, subStr string) bool {
	for _, c := range subStr {
		if !strings.ContainsRune(s, c) {
			return false
		}
	}
	return true
}

func numContainingRunes(s, subStr string) int {
	count := 0
	for _, c := range subStr {
		if strings.ContainsRune(s, c) {
			count++
		}
	}
	return count
}

func getMapping(inputSegments []string) [10]string {
	numberMapping := [10]string{}
	for _, segment := range inputSegments {
		switch len(segment) {
		case 2:
			numberMapping[1] = segment
		case 3:
			numberMapping[7] = segment
		case 4:
			numberMapping[4] = segment
		case 7:
			numberMapping[8] = segment
		}
	}
	for _, segment := range inputSegments {
		if len(segment) == 5 {
			if containsAllRunes(segment, numberMapping[1]) {
				numberMapping[3] = segment
			} else if numContainingRunes(segment, numberMapping[4]) == 3 {
				numberMapping[5] = segment
			} else {
				numberMapping[2] = segment
			}
		} else if len(segment) == 6 {
			if containsAllRunes(segment, numberMapping[4]) {
				numberMapping[9] = segment
			} else if containsAllRunes(segment, numberMapping[7]) {
				numberMapping[0] = segment
			} else {
				numberMapping[6] = segment
			}
		}
	}
	return numberMapping
}

func main() {
	lines := readLines("input.txt")
	total := 0
	for _, line := range lines {
		parts := strings.Split(line, " | ")
		inputSegments := strings.Split(parts[0], " ")
		outputSegments := strings.Split(parts[1], " ")
		mapping := getMapping(inputSegments)
		number := 0
		for _, outputSegment := range outputSegments {
			for i, inputSegment := range mapping {
				if containsAllRunes(outputSegment, inputSegment) && containsAllRunes(inputSegment, outputSegment) {
					number *= 10
					number += i
					break
				}
			}
		}
		total += number
	}
	fmt.Println(total)
}
