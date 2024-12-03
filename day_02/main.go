package main

import (
	"aoc_2024/utils"
	"strconv"
	"strings"
)

func main() {
	input, part := utils.ParseInput(2)
	if part == 1 {
		part_1(input)
	} else {
		part_2(input)
	}
}

func part_1(input string) {
	reports := parseReports(input)
	safeReportCount := 0
	for _, report := range reports {
		if isReportSafe(report, -1) {
			safeReportCount++
		}
	}

	println(safeReportCount)
}

func part_2(input string) {
	reports := parseReports(input)
	safeReportCount := 0
	for _, report := range reports {
		if isReportSafe(report, -1) {
			safeReportCount++
			continue
		}
		for skipIndex := range report {
			if isReportSafe(report, skipIndex) {
				safeReportCount++
				break
			}
		}
	}

	println(safeReportCount)
}

func isReportSafe(report Report, skip int) bool {
	startIndex := 0
	if skip == 0 {
		startIndex = 1
	}
	previousLevel := report[startIndex]
	direction := 0
	isSafe := true
	for j, level := range report {
		if j == startIndex || j == skip  {
			continue
		}

		if direction == 0 {
			if level > previousLevel {
				direction = 1
			} else {
				direction = -1
			}
		} else {
			if (level > previousLevel) != (direction == 1) {
				// UNSAFE via direction change
				isSafe = false
				previousLevel = level
				break
			}
		}

		if (level - previousLevel) * direction < 1 || (level - previousLevel) * direction > 3 {
			// UNSAFE via excessive change
			isSafe = false
			previousLevel = level
			break
		}
		previousLevel = level
	}
	return isSafe
}

type Report []int

func parseReports(input string) []Report {
	var reports []Report
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			break
		}
		numbers := strings.Split(line, " ")
		var report Report
		for _, numStr := range numbers {
			num, _ := strconv.Atoi(numStr)
			report = append(report, num)
		}
		reports = append(reports, report)
	}
	return reports
}
