package main

import (
	"aoc_2024/utils"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, part := utils.ParseInput(1)
	if part == 1 {
		part_1(input)
	} else {
		part_2(input)
	}
}

func part_1(input string) {
	list1, list2 := parseLists(input)
	slices.Sort(list1)
	slices.Sort(list2)

	var diffSum int
	for i := range list1 {
		if list1[i] > list2[i] {
			diffSum += list1[i] - list2[i]
		} else {
			diffSum += list2[i] - list1[i]
		}
	}

	println(diffSum)
}

func part_2(input string) {
	numberFrequencies := make(map[int]int)

	list1, list2 := parseLists(input)
	for _, num := range list2 {
		numberFrequencies[num]++
	}

	var similarityScore int

	for _, num := range list1 {
		similarityScore += num * numberFrequencies[num]
	}

	println(similarityScore)
}

func parseLists(input string) ([]int, []int) {
	var list1 []int
	var list2 []int
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			break
		}
		numbers := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	return list1, list2
}
