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

func getLists(filePath string) ([]int, []int) {
	inputFile, _ := os.Open(filePath)
	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}(inputFile)
	var leftList []int
	var rightList []int

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

	return leftList, rightList
}

func findDistance() int {
	var finalDistance int

	leftList, rightList := getLists("input.txt")

	for i, left := range leftList {
		finalDistance += int(math.Abs(float64(left - rightList[i])))
	}

	return finalDistance
}

func findSimilarityScore() int {
	leftList, rightList := getLists("input.txt")
	sum := 0
	hashMap := make(map[int]int)

	for _, item := range rightList {
		// if value is in hash map, add its count
		if _, exists := hashMap[item]; exists {
			hashMap[item] += 1
		} else {
			hashMap[item] = 1
		}
	}

	for _, v := range leftList {
		if _, exists := hashMap[v]; exists {
			sum += v * hashMap[v]
		}
	}

	return sum
}

func main() {
	fmt.Println(findDistance())
	fmt.Println(findSimilarityScore())
}
