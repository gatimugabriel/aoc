package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var mulPattern = `mul\((\d{1,3}),(\d{1,3})\)`

func getInput(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error opening file", err)
	}

	return string(content)
}

// part 1
func getMulValues(input string) int {
	sum := 0
	mulRegex := regexp.MustCompile(mulPattern)
	matches := mulRegex.FindAllStringSubmatch(input, -1)

	for _, mul := range matches {
		x, _ := strconv.Atoi(mul[1])
		y, _ := strconv.Atoi(mul[2])
		sum += x * y
	}

	return sum
}

// part 2
func getEnabledValues(input string) int {
	sum := 0
	do := true

	doRegex := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	matches := doRegex.FindAllString(input, -1)

	for _, instruction := range matches {
		//fmt.Println("instruction", instruction)

		switch {
		case instruction == "do()":
			do = true
		case do && regexp.MustCompile(mulPattern).MatchString(instruction):
			sum += getMulValues(instruction)
		default:
			do = false
		}
	}
	return sum
}

func main() {
	input := getInput("input.txt")
	fmt.Println("Part1: ", getMulValues(input))
	fmt.Println("Part2: ", getEnabledValues(input))
}
