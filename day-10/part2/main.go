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

type RuneStack []rune

func (runeStack *RuneStack) push(r rune) {
	*runeStack = append(*runeStack, r)
}

func (runeStack *RuneStack) pop() rune {
	r := (*runeStack)[len(*runeStack)-1]
	*runeStack = (*runeStack)[:len(*runeStack)-1]
	return r
}

func (runeStack *RuneStack) isEmpty() bool {
	return len(*runeStack) == 0
}

func isClosing(a rune) bool {
	return a == ')' || a == ']' || a == '}' || a == '>'
}

func areMatching(a, b rune) bool {
	return a == '(' && b == ')' || a == '[' && b == ']' || a == '{' && b == '}' || a == '<' && b == '>'
}

func closing(a rune) rune {
	switch a {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	default:
		panic("invalid rune")
	}
}

func main() {
	lines := readLines("input.txt")
	scoreTable := map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
	var scores []int
	for _, line := range lines {
		var runeStack RuneStack
		incomplete := true
		for _, c := range line {
			if isClosing(c) {
				if !areMatching(runeStack.pop(), c) {
					incomplete = false
					break
				}
			} else {
				runeStack.push(c)
			}
		}
		if incomplete {
			score := 0
			for !runeStack.isEmpty() {
				c := runeStack.pop()
				score *= 5
				score += scoreTable[closing(c)]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
