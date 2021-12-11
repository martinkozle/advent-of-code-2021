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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	lines := readLines("input.txt")
	var positions []int
	max := 0
	for _, positionString := range strings.Split(lines[0], ",") {
		position := toInt(positionString)
		positions = append(positions, position)
		if position > max {
			max = position
		}
	}
	lowestCost := max * len(positions)
	for goal := 0; goal <= max; goal++ {
		cost := 0
		for _, position := range positions {
			cost += abs(goal - position)
		}
		if cost < lowestCost {
			lowestCost = cost
		}
	}
	fmt.Println(lowestCost)
}
