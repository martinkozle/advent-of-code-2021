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
	distance := 0
	depth := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		amount := toInt(parts[1])
		if direction == "forward" {
			distance += amount
		} else if direction == "down" {
			depth += amount
		} else if direction == "up" {
			depth -= amount
		}
	}
	fmt.Println(distance * depth)
}
