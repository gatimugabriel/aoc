package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputFile, _ := os.Open("input.txt")
	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {

		}
	}(inputFile)

	var leftList []int
	var rightList []int
	var finalDistance int

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)

		leftInt, _ := strconv.Atoi(parts[0])
		rightInt, _ := strconv.Atoi(parts[1])

		leftList = append(leftList, leftInt)
		rightList = append(rightList, rightInt)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	for i, left := range leftList {
		finalDistance += int(math.Abs(float64(left - rightList[i])))
	}

	fmt.Println(finalDistance)
}
