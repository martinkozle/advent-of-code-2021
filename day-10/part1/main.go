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

type RuneStack []rune

func (runeStack *RuneStack) push(r rune) {
	*runeStack = append(*runeStack, r)
}

func (runeStack *RuneStack) pop() rune {
	r := (*runeStack)[len(*runeStack)-1]
	*runeStack = (*runeStack)[:len(*runeStack)-1]
	return r
}

func (runeStack RuneStack) isEmpty() bool {
	return len(runeStack) == 0
}

func isClosing(a rune) bool {
	return a == ')' || a == ']' || a == '}' || a == '>'
}

func areMatching(a, b rune) bool {
	return a == '(' && b == ')' || a == '[' && b == ']' || a == '{' && b == '}' || a == '<' && b == '>'
}

func main() {
	lines := readLines("input.txt")
	scoreTable := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	totalScore := 0
	for _, line := range lines {
		var runeStack RuneStack
		for _, c := range line {
			if isClosing(c) {
				score := scoreTable[c]
				if !areMatching(runeStack.pop(), c) {
					totalScore += score
					break
				}
			} else {
				runeStack.push(c)
			}
		}
	}
	fmt.Println(totalScore)
}
