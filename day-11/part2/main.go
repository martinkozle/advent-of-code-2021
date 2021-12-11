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

type OctopusGrid [10][10]int

func inRange(y, x int) bool {
	if x < 0 || x > 9 || y < 0 || y > 9 {
		return false
	}
	return true
}

func (grid *OctopusGrid) flash(y, x int) {
	for oy := -1; oy <= 1; oy++ {
		for ox := -1; ox <= 1; ox++ {
			if oy == 0 && ox == 0 || !inRange(y+oy, x+ox) {
				continue
			}
			if grid[y+oy][x+ox] != -1 {
				grid[y+oy][x+ox]++
				if grid[y+oy][x+ox] > 9 {
					grid[y+oy][x+ox] = -1
					grid.flash(y+oy, x+ox)
				}
			}
		}
	}
}

func (grid *OctopusGrid) step() int {
	count := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			grid[i][j]++
		}
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if grid[i][j] > 9 {
				grid[i][j] = -1
				grid.flash(i, j)
			}
		}
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if grid[i][j] == -1 {
				count++
				grid[i][j] = 0
			}
		}
	}
	return count
}

func main() {
	lines := readLines("input.txt")
	var grid OctopusGrid
	for i, line := range lines {
		for j, c := range line {
			grid[i][j] = toInt(string(c))
		}
	}
	for i := 0; true; i++ {
		count := grid.step()
		if count == 100 {
			fmt.Println(i + 1)
			break
		}
	}
}
