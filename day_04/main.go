package main

import (
	"aoc_2024/utils"
	"fmt"
	"strings"
)

func main() {
	fmt.Print("")// Stop warning me about an unused "fmt" package!
	input, part := utils.ParseInput(2)
	if part == 1 {
		part_1(input)
	} else {
		part_2(input)
	}
}

type Grid [][]rune
type Vector struct {
	x int
	y int
}
const Happy = true
const Sad = false

func part_1(input string) {
	keyword := "XMAS"
	var wordCount int

	grid := parseInputIntoGrid(input)
	for y := range grid {
		for x := range grid[y] {
			gridCell := grid[y][x]
			if gridCell == rune(keyword[0]) {
				wordCount += allDirectionWordScan(grid, x, y, keyword)
			}
		}
	}

	println(wordCount)
}

func addVectors(p1 Vector, p2 Vector) Vector {
	return Vector {
		x: p1.x + p2.x,
		y: p1.y + p2.y,
	}
}

func invertVector(p Vector) Vector {
	return Vector {
		x: -p.x,
		y: -p.y,
	}
}

func (v *Vector) isInBounds(x int, y int) bool {
	return v.x >= 0 && v.x < x && v.y >= 0 && v.y < y
}

func parseInputIntoGrid(input string) Grid {
	grid := make(Grid, 0)
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		if len(line) == 0 {
			break
		}
		grid = append(grid, []rune{})
		for _, char := range line {
			grid[y] = append(grid[y], rune(char))
		}
	}
	return grid
}

func allDirectionWordScan(grid Grid, x int, y int, keyword string) int {
	maxY := len(grid)
	maxX := len(grid[0])
	var checkDirections []Vector

	// sanity
	if grid[y][x] != rune(keyword[0]) {
		return 0
	}

	//   -1
	// -1   1
	//    1
	checkDirections = append(checkDirections, Vector{x: 0, y: -1})
	checkDirections = append(checkDirections, Vector{x: 0, y: 1})
	checkDirections = append(checkDirections, Vector{x: -1, y: 0})
	checkDirections = append(checkDirections, Vector{x: 1, y: 0})
	checkDirections = append(checkDirections, Vector{x: -1, y: -1})
	checkDirections = append(checkDirections, Vector{x: 1, y: -1})
	checkDirections = append(checkDirections, Vector{x: -1, y: 1})
	checkDirections = append(checkDirections, Vector{x: 1, y: 1})


	var foundWordCount int
	for _, direction := range checkDirections {
		currentPosition := Vector{x: x, y: y}
		check := Happy
		for _, char := range keyword {
			// Check bounds and letter in word
			if !currentPosition.isInBounds(maxX, maxY) || grid[currentPosition.y][currentPosition.x] != rune(char) {
				check = Sad
				break
			}
			currentPosition = addVectors(currentPosition, direction)
		}
		if check == Happy {
			foundWordCount++
		}
	}

	return foundWordCount
}

func part_2(input string) {
	keyword := "MAS"
	middleIndex := len(keyword) / 2
	var crossingCount int

	grid := parseInputIntoGrid(input)
	for y := range grid {
		for x := range grid[y] {
			gridCell := grid[y][x]
			if gridCell == rune(keyword[middleIndex]) {
				if isWordCross(grid, x, y, keyword) {
					crossingCount++
				}
			}
		}
	}

	println(crossingCount)
}

func isWordCross(grid Grid, x int, y int, keyword string) bool {
	if len(keyword) & 1 == 0 {
		panic("Keyword is not an odd length")
	}

	middleIndex := len(keyword) / 2
	maxY := len(grid)
	maxX := len(grid[0])
	var checkDirections []Vector

	// sanity
	if grid[y][x] != rune(keyword[middleIndex]) {
		panic("position is not middle index of keyword")
	}

	//   -1
	// -1   1
	//    1
	// Order here matters to work with previousCheck
	checkDirections = append(checkDirections, Vector{x: -1, y: -1})
	checkDirections = append(checkDirections, Vector{x: 1, y: -1})
	checkDirections = append(checkDirections, Vector{x: 1, y: 1})
	checkDirections = append(checkDirections, Vector{x: -1, y: 1})

	previousCheck := Sad
	for _, direction := range checkDirections {
		positivePosition := Vector{x: x, y: y}
		negativePosition := Vector{x: x, y: y}
		check := Happy

		for offset := range middleIndex + 1 {
			if  !positivePosition.isInBounds(maxX, maxY) ||
				!negativePosition.isInBounds(maxX, maxY) ||
				grid[positivePosition.y][positivePosition.x] != rune(keyword[middleIndex + offset]) ||
				grid[negativePosition.y][negativePosition.x] != rune(keyword[middleIndex - offset]) {

				check = Sad
				break
			}

			positivePosition = addVectors(positivePosition, direction)
			negativePosition = addVectors(negativePosition, invertVector(direction))
		}
		if check == Happy {
			if previousCheck == Happy {
				return true
			}
			previousCheck = Happy
		}
	}
	return false

}
