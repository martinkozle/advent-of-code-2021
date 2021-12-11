package main

import (
	"fmt"
	"os"
	"sort"
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

func (heightMap *HeightMap) get(x, y int) int {
	if x < 0 || y < 0 || x >= 100 || y >= 100 {
		return 9
	}
	return heightMap[y][x]
}

type Visited [100][100]bool

func (visited *Visited) get(x, y int) bool {
	if x < 0 || y < 0 || x >= 100 || y >= 100 {
		return true
	}
	return visited[y][x]
}

type Coordinate struct {
	x, y int
}

func floodFill(heightMap *HeightMap, visited *Visited, x, y int) int {
	if visited.get(x, y) {
		return 0
	}
	visited[y][x] = true
	offsets := []Coordinate{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	count := 1
	for _, offset := range offsets {
		count += floodFill(heightMap, visited, x+offset.x, y+offset.y)
	}
	return count
}

func main() {
	lines := readLines("input.txt")
	var heightMap HeightMap
	var visited Visited

	for i, line := range lines {
		for j, c := range line {
			heightMap[i][j] = toInt(string(c))
			if heightMap[i][j] == 9 {
				visited[i][j] = true
			}
		}
	}
	var counts []int
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			count := floodFill(&heightMap, &visited, i, j)
			if count > 0 {
				counts = append(counts, count)
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	fmt.Println(counts[0] * counts[1] * counts[2])
}
