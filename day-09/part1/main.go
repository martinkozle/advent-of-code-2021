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

type HeightMap [100][100]int

type Coordinate struct {
	x, y int
}

func (heightMap *HeightMap) get(x, y int) int {
	if x < 0 || y < 0 || x >= 100 || y >= 100 {
		return 10
	}
	return (*heightMap)[y][x]
}

func main() {
	lines := readLines("input.txt")
	var heightMap HeightMap
	for i, line := range lines {
		for j, c := range line {
			heightMap[i][j] = toInt(string(c))
		}
	}
	offsets := []Coordinate{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	sum := 0
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			flag := true
			current := heightMap.get(j, i)
			for _, offset := range offsets {
				adjacent := heightMap.get(j+offset.x, i+offset.y)
				if current >= adjacent {
					flag = false
					break
				}
			}
			if flag {
				sum += current + 1
			}
		}
	}
	fmt.Println(sum)
}
