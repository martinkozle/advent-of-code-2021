package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)
	commands := strings.Split(string(data), "\n")
	distance := 0
	aim := 0
	depth := 0
	for _, command := range commands {
		if command == "" {
			continue
		}
		parts := strings.Split(command, " ")
		direction := parts[0]
		amount, err := strconv.Atoi(parts[1])
		check(err)
		if direction == "forward" {
			distance += amount
			depth += aim * amount
		} else if direction == "down" {
			aim += amount
		} else if direction == "up" {
			aim -= amount
		}
	}
	fmt.Println(distance * depth)
}
