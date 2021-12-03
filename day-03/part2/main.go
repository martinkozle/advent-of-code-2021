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

func countBits(readings []string) []int {
	gammaBits := make([]int, len(readings[0]))
	for _, reading := range readings {
		for i, bit := range reading {
			if bit == '1' {
				gammaBits[i]++
			} else {
				gammaBits[i]--
			}
		}
	}
	return gammaBits
}

func main() {
	lines := readLines("input.txt")
	bitLen := len(lines[0])
	o2Readings := append(make([]string, 0, len(lines)), lines...)
	for i := 0; i < bitLen && len(o2Readings) > 1; i++ {
		var temp []string
		gammaBits := countBits(o2Readings)
		for _, reading := range o2Readings {
			if (gammaBits[i] >= 0 && reading[i] == '1') || (gammaBits[i] < 0 && reading[i] == '0') {
				temp = append(temp, reading)
			}
		}
		o2Readings = append(o2Readings[:0], temp...)
	}
	o2Bits := o2Readings[0]
	co2Readings := append(make([]string, 0, len(lines)), lines...)
	for i := 0; i < bitLen && len(co2Readings) > 1; i++ {
		var temp []string
		gammaBits := countBits(co2Readings)
		for _, reading := range co2Readings {
			if (gammaBits[i] >= 0 && reading[i] == '0') || (gammaBits[i] < 0 && reading[i] == '1') {
				temp = append(temp, reading)
			}
		}
		co2Readings = append(co2Readings[:0], temp...)
	}
	co2Bits := co2Readings[0]
	o2 := 0
	co2 := 0
	for i := 0; i < bitLen; i++ {
		if o2Bits[i] == '1' {
			o2 |= 1 << (bitLen - i - 1)
		}
		if co2Bits[i] == '1' {
			co2 |= 1 << (bitLen - i - 1)
		}
	}
	fmt.Println(o2 * co2)
}
