package main

import (
	"aoc_2024/utils"
	"aoc_2024/utils/vec"
	"fmt"
	"strconv"
)

func main() {
	fmt.Print()
	input, part := utils.ParseInput(2)
	if part == 1 {
		part_1(input)
	} else {
		part_2(input)
	}
}

type TrailPoint struct {
	v vec.Vec
	h int
}

const MAX_HEIGHT = 9

func part_1(input string) {
	// What am I even doing? This does not answer the question at all...
	vm := vec.ParseIntoMap(input)
	validPaths := make([]TrailPoint, 0)
	for v, char := range vm.Vm {
		h, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}

		if h == 0 {
			validPaths = append(validPaths, TrailPoint{v: v, h: h})
		}
	}

	nextPaths := make([]TrailPoint, 0)
	finalPathPoints := make(map[vec.Vec]struct{}, 0)

	for len(validPaths) > 0 {
		for _, tp := range validPaths {
			for _, o := range vm.NeighborsTo(tp.v) {
				otherChar := vm.Vm[o]
				// fmt.Printf("\n%v\n", o)
				// vm.Print()
				h, err := strconv.Atoi(string(otherChar))
				if err != nil {
					panic(err)
				}

				if (h - tp.h) == 1 {
					if h == MAX_HEIGHT {
						fmt.Printf("Max Found at %v\n", o)
						finalPathPoints[o] = struct{}{}
					}
					nextPaths = append(nextPaths, TrailPoint{v: o, h: h})
				}
			}
		}
		validPaths = nextPaths
		nextPaths = make([]TrailPoint, 0)
	}
	
	fmt.Printf("%d\n", len(finalPathPoints))
}

func part_2(input string) {
	println("Part 2 START")
	println(input)
	println("Part 2 END")
}
