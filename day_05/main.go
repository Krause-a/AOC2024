package main

import (
	"aoc_2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("")//Use fmt package to make development easier
	input, part := utils.ParseInput(2)
	if part == 1 {
		part_1(input)
	} else {
		part_2(input)
	}
}

func part_1(input string) {
	orderings, updates := parseInput(input)
	var correctUpdates []Update

	for _, update := range updates {
		if update.isOrderingCorrect(orderings) {
			correctUpdates = append(correctUpdates, update)
		}
	}

	sum := updatesMiddleSum(correctUpdates)

	println(sum)
}

const Happy = true
const Sad = false
type Page int
type PageOrdering struct {
	isSatisfied bool
	first Page
	second Page
}
type Update []Page

func parseInput(input string) ([]PageOrdering, []Update) {
	lines := strings.Split(input, "\n")
	onOrderings := true

	var orderings []PageOrdering
	var updates []Update

	for _, line := range lines {
		if len(line) == 0 {
			if onOrderings {
				onOrderings = false
				continue
			}
			break
		}
		if onOrderings {
			orderingsSplit := strings.Split(line, "|")
			firstPage, _ := strconv.Atoi(orderingsSplit[0])
			secondPage, _ := strconv.Atoi(orderingsSplit[1])

			orderings = append(orderings, PageOrdering{
				isSatisfied: false,
				first: Page(firstPage),
				second: Page(secondPage),
			})
		} else {
			pagesSplit := strings.Split(line, ",")
			var update Update
			for _, pageStr := range pagesSplit {
				page, _ := strconv.Atoi(pageStr)
				update = append(update, Page(page))
			}
			updates = append(updates, update)
		}
	}
	return orderings, updates
}

func resetOrderings(orderings *[]PageOrdering) {
	for i := range *orderings {
		ordering := &(*orderings)[i]
		ordering.isSatisfied = false
	}
}

func updateContainsOrdering(update *Update, ordering *PageOrdering) bool {
	containsFirst := false
	containsSecond := false
	for _, page := range *update {
		if page == ordering.first {
			containsFirst = true
		}
		if page == ordering.second {
			containsSecond = true
		}
		if containsFirst && containsSecond {
			break
		}
	}
	return containsFirst && containsSecond
}

func updatesMiddleSum(updates []Update) int {
	var sum int
	for _, update := range updates {
		sum += int(update[len(update)/2])
	}
	return sum
}

func part_2(input string) {
	orderings, updates := parseInput(input)

	var incorrectUpdates []Update

	for _, update := range updates {
		if !update.isOrderingCorrect(orderings) {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}

	var middleSum int

	for len(incorrectUpdates) > 0 {
		var nextUpdates []Update
		for _, update := range incorrectUpdates {
			correctedUpdate := correctUpdateOrdering(update, orderings)
			if correctedUpdate.isOrderingCorrect(orderings) {
				middleSum += int(correctedUpdate[len(correctedUpdate)/2])
			} else {
				nextUpdates = append(nextUpdates, correctedUpdate)
			}
		}
		incorrectUpdates = nextUpdates
	}

	println(middleSum)
}

type UpdateCorrection struct {
	update Update
	corrections []Correction
}

type Correction struct {
	aIndex int
	bIndex int
}

func (u *Update) isOrderingCorrect(orderings []PageOrdering) bool {
	check := Happy
	for i := range orderings {
		ordering := &orderings[i]
		if !updateContainsOrdering(u, ordering) {
			continue
		}
		for _, page := range *u {
			if page == ordering.first {
				ordering.isSatisfied = true
			}
			if page == ordering.second && !ordering.isSatisfied {
				check = Sad
				break
			}
		}
		if (check == Sad) {
			break
		}
	}
	resetOrderings(&orderings)
	return check == Happy
}

func correctUpdateOrdering(update Update, orderings []PageOrdering) Update {
	check := Happy
	updateCorrection := UpdateCorrection{
		update: update,
		corrections: []Correction{},
	}
	for i := range orderings {
		ordering := &orderings[i]
		if !updateContainsOrdering(&update, ordering) {
			continue
		}
		incorrectIndex := -1
		for pageIndex, page := range update {
			if page == ordering.first {
				ordering.isSatisfied = true
			}
			if page == ordering.second && !ordering.isSatisfied {
				check = Sad
				incorrectIndex = pageIndex
				break
			}
		}
		if (check == Sad) {
			if (incorrectIndex == -1) {
				break
			}
			for pageIndex, page := range update {
				if page == ordering.first {
					updateCorrection.corrections = append(updateCorrection.corrections, Correction{
						aIndex: pageIndex,
						bIndex: incorrectIndex,
					})
					break
				}
			}
		}
	}
	resetOrderings(&orderings)
	var swappedIndices []int
	update = updateCorrection.update
	corrections := updateCorrection.corrections
	for _, correction := range corrections {
		if utils.Contains(swappedIndices, correction.aIndex) || utils.Contains(swappedIndices, correction.bIndex) {
			continue
		}
		swappedIndices = append(swappedIndices, correction.aIndex)
		swappedIndices = append(swappedIndices, correction.bIndex)
		temp := update[correction.bIndex]
		update[correction.bIndex] = update[correction.aIndex]
		update[correction.aIndex] = temp
	}
	return update
}
