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

	counts := [9]int{}
	for _, timerString := range strings.Split(lines[0], ",") {
		timer := toInt(timerString)
		counts[timer]++
	}

	for day := 0; day < 256; day++ {
		temp := [9]int{}
		for timer := 1; timer <= 8; timer++ {
			temp[timer-1] = counts[timer]
		}
		temp[6] += counts[0]
		temp[8] += counts[0]
		counts = temp
	}

	total := 0
	for _, count := range counts {
		total += count
	}

	fmt.Println(total)
}
