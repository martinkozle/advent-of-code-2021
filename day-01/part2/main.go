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
	depthSums := []int{}
	for i := 2; i < len(lines); i++ {
		depth1 := toInt(lines[i-2])
		depth2 := toInt(lines[i-1])
		depth3 := toInt(lines[i])
		depthSums = append(depthSums, depth1+depth2+depth3)
	}
	count := 0
	for i := 1; i < len(depthSums); i++ {
		if depthSums[i] > depthSums[i-1] {
			count++
		}
	}
	fmt.Println(count)
}
