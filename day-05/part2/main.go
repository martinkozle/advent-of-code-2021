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

type Point struct {
	x, y int
}

type LineSegment struct {
	p1, p2 Point
}

type Field struct {
	rows [1000][1000]int
}

func (field *Field) numberOfIntersections() int {
	count := 0
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			if field.rows[y][x] > 1 {
				count++
			}
		}
	}
	return count
}

func main() {
	lines := readLines("input.txt")
	lineSegments := make([]LineSegment, len(lines))
	for i, line := range lines {
		lineParts := strings.Split(line, " -> ")
		point1Parts := strings.Split(lineParts[0], ",")
		point2Parts := strings.Split(lineParts[1], ",")
		x1 := toInt(point1Parts[0])
		y1 := toInt(point1Parts[1])
		x2 := toInt(point2Parts[0])
		y2 := toInt(point2Parts[1])
		lineSegments[i] = LineSegment{Point{x1, y1}, Point{x2, y2}}
	}
	var field Field
	for _, lineSegment := range lineSegments {
		startX := lineSegment.p1.x
		endX := lineSegment.p2.x
		startY := lineSegment.p1.y
		endY := lineSegment.p2.y

		skipX := 1
		skipY := 1

		if startX == endX {
			skipX = 0
		} else if startX > endX {
			skipX = -1
		}

		if startY == endY {
			skipY = 0
		} else if startY > endY {
			skipY = -1
		}

		x, y := startX, startY

		for x != endX || y != endY {
			field.rows[y][x]++
			x += skipX
			y += skipY
		}

		field.rows[y][x]++
	}

	fmt.Println(field.numberOfIntersections())

}
