package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func getInput(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error opening file", err)
	}

	return string(content)
}

func main() {
	input := getInput("input.txt")

	mulPattern := `mul\((\d{1,3}),(\d{1,3})\)`
	mulRegex := regexp.MustCompile(mulPattern)
	matches := mulRegex.FindAllStringSubmatch(input, -1)

	sum := 0

	for _, mul := range matches {
		x, _ := strconv.Atoi(mul[1])
		y, _ := strconv.Atoi(mul[2])
		sum += x * y
	}

	fmt.Println(sum)
}
