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

type Board struct {
	rows [][]int
}

type Boards struct {
	boards []*Board
}

func readBoard(lines []string) *Board {
	rows := make([][]int, len(lines))
	for i, line := range lines {
		split := strings.Fields(line)
		row := make([]int, len(split))
		for j, element := range split {
			row[j] = toInt(element)
		}
		rows[i] = row
	}
	return &Board{rows}
}

func readBoards(lines []string) Boards {
	var boards Boards
	for i := 2; i < 100*6; i += 6 {
		boards.boards = append(boards.boards, readBoard(lines[i:i+5]))
	}
	return boards
}

func (board *Board) markBoard(number int) {
	for i, row := range board.rows {
		for j, element := range row {
			if element == number {
				board.rows[i][j] = -1
			}
		}
	}
}

func (board *Board) checkBoard() bool {
	for i := range board.rows {
		flag := true
		for j := range board.rows[0] {
			if board.rows[i][j] != -1 {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	for i := range board.rows[0] {
		flag := true
		for j := range board.rows {
			if board.rows[j][i] != -1 {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	return false
}

func (board *Board) sumBoard() int {
	sum := 0
	for _, row := range board.rows {
		for _, element := range row {
			if element != -1 {
				sum += element
			}
		}
	}
	return sum
}

func main() {
	lines := readLines("input.txt")
	numberStrings := strings.Split(lines[0], ",")
	numbers := make([]int, len(numberStrings))
	for i, numberString := range numberStrings {
		numbers[i] = toInt(numberString)
	}
	boards := readBoards(lines)
	for _, number := range numbers {
		for _, board := range boards.boards {
			board.markBoard(number)
			if board.checkBoard() {
				fmt.Println(board.sumBoard() * number)
				return
			}
		}
	}
}
