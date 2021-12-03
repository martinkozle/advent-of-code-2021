package main

import (
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readLines(file string) []string {
	data, err := os.ReadFile(file)
	check(err)
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func main() {
	lines := readLines("input.txt")
	bitLen := len(lines[0])
	gammaBits := make([]int, bitLen)
	for _, line := range lines {
		for i, bit := range line {
			if bit == '1' {
				gammaBits[i]++
			} else {
				gammaBits[i]--
			}
		}
	}
	gamma := 0
	epsilon := 0
	for i, count := range gammaBits {
		if count > 0 {
			gamma |= 1 << (bitLen - i - 1)
		} else {
			epsilon |= 1 << (bitLen - i - 1)
		}
	}
	fmt.Println(gamma * epsilon)
}
