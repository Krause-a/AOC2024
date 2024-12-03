package main

import (
	"aoc_2024/utils"
	"regexp"
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
	validPattern := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	numberPattern := regexp.MustCompile(`\d{1,3}`)
	validStatements := validPattern.FindAllString(input, -1)
	var total int
	for _, statement := range validStatements {
		numbers := numberPattern.FindAllString(statement, -1)
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		total += num1 * num2
	}

	println(total)
}

func part_2(input string) {
	input = strings.ReplaceAll(input, "\n"," ")
	doPattern := regexp.MustCompile(`(^|do\(\)).*?(don't\(\)|\n|$)`)
	validPattern := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	numberPattern := regexp.MustCompile(`\d{1,3}`)
	doSections := doPattern.FindAllString(input, -1)
	var total int
	for _, section := range doSections {
		validStatements := validPattern.FindAllString(section, -1)
		for _, statement := range validStatements {
			numbers := numberPattern.FindAllString(statement, -1)
			num1, _ := strconv.Atoi(numbers[0])
			num2, _ := strconv.Atoi(numbers[1])
			total += num1 * num2
		}
	}

	println(total)
}
