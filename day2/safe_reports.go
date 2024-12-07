package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getReports(filePath string) [][]int {
	inputFile, _ := os.Open(filePath)
	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}(inputFile)

	var reports [][]int

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		report := scanner.Text()                // each line is a list(report)
		reportStrings := strings.Fields(report) // each part in the report(line) is a level

		reportLevels := make([]int, len(reportStrings))
		for i, reportLevel := range reportStrings {
			level, _ := strconv.Atoi(reportLevel)
			reportLevels[i] = level
		}

		reports = append(reports, reportLevels)
	}

	return reports
}

func isReportSafe(report []int) bool {
	isSafe := true
	//  check direction (increasing/decreasing)
	isIncreasing := report[0] < report[1]

	for i := 0; i < len(report)-1; i++ {
		a := report[i]
		b := report[i+1]

		if isIncreasing {
			if a > b || a < b-3 || a == b {
				isSafe = false
				break
			}
		} else {
			if a < b || a > b+3 || a == b {
				isSafe = false
				break
			}
		}
	}

	return isSafe
}

// safeReports; Part 1
func safeReports(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		isSafe := isReportSafe(report)
		if isSafe {
			safeReports++
		}
	}

	return safeReports
}

// problemDampener: part 2
func problemDampener(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		isSafe := isReportSafe(report)
		if isSafe {
			safeReports++
		} else {
			// remove one level and check again
			for i := 0; i < len(report); i++ {
				modifiedReport := make([]int, 0, len(report)-1)
				modifiedReport = append(modifiedReport, report[:i]...)
				modifiedReport = append(modifiedReport, report[i+1:]...)
			
				if isReportSafe(modifiedReport) {
					// fmt.Printf("report %d is now safe after removing level %d\t => %v\n\n", x+1, i+1, modifiedReport)
					safeReports += 1
					break
				} else{
					// fmt.Printf("report number %d  => %v is still unsafe \t original : %v\n\n", x+1, modifiedReport, report)
				}
			}
		}
	} 

	return safeReports
}

func main() {
	reports := getReports("input.txt")
	fmt.Println("Safe reports = ", safeReports(reports))
	fmt.Println("Safe reports after using problem dampener = ", problemDampener(reports))
}
