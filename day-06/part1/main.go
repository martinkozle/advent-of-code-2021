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

type Lanternfish struct {
	timer int
}

func main() {
	lines := readLines("input.txt")

	var fishes []Lanternfish
	for _, timerString := range strings.Split(lines[0], ",") {
		timer := toInt(timerString)
		fishes = append(fishes, Lanternfish{timer})
	}
	for day := 0; day < 80; day++ {
		for i := range fishes {
			fishes[i].timer--
			if fishes[i].timer == -1 {
				fishes[i].timer = 6
				fishes = append(fishes, Lanternfish{8})
			}
		}
	}

	fmt.Println(len(fishes))
}
