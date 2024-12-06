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

func main() {
	reports := getReports("input.txt")
	safeReports := 0

	for _, report := range reports {
		// fmt.Printf("Checking report number %d: %v\n", reportNumber+1, report)

		//  check direction (either increasing/decreasing)
		isIncreasing := report[0] < report[1]
		// if isIncreasing {
		// 	fmt.Printf("Report number %d is increasing\n", reportNumber+1)
		// } else {
		// 	fmt.Printf("Report number %d is decreasing\n", reportNumber+1)
		// }

		// determine whether it is safe
		isSafeReport := true
		for i := 0; i < len(report)-1; i++ {
			if isIncreasing {
				if report[i] > report[i+1] || report[i] < report[i+1]-3 || report[i] == report[i+1] {
					// fmt.Printf("Breaking out of report  %d: %d > %d\n", reportNumber, report[i], report[i+1])
					isSafeReport = false
					break
				}
			} else {
				if report[i] < report[i+1] || report[i] > report[i+1]+3 || report[i] == report[i+1]{
					// fmt.Printf("Breaking out of report  %d: %d > %d\n", reportNumber, report[i], report[i+1])
					isSafeReport = false
					break
				}
			}
		}

		if isSafeReport {
			// fmt.Printf("Report number %d is safe\n", reportNumber+1)
			safeReports++
		}
	}

	fmt.Println(safeReports)
}
