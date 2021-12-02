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

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	count := 0
	for i := 1; i < len(lines); i++ {
		depth1 := toInt(lines[i-1])
		depth2 := toInt(lines[i])
		if depth1 < depth2 {
			count++
		}
	}
	fmt.Println(count)
}
