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

func main() {
	lines := readLines("input.txt")
	countedLengths := map[int]bool{2: true, 4: true, 3: true, 7: true}
	count := 0
	for _, line := range lines {
		parts := strings.Split(line, " | ")
		outputSegments := strings.Split(parts[1], " ")
		for _, segment := range outputSegments {
			if countedLengths[len(segment)] {
				count++
			}
		}
	}
	fmt.Println(count)
}
