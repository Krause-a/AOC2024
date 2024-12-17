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

	sum := 0
	otherSum := 0
	for _, v := range validPaths {
		foundTrailends := make(map[vec.Vec]struct{})
		amount := StepAll(&foundTrailends, &vm, v.v)
		otherSum += amount
		// fmt.Printf("Amount for %v == %d,,,%d\n", v.v, amount, len(foundTrailends))
		sum += len(foundTrailends)
	}

	fmt.Printf("Part 1: %d\n", sum)
	fmt.Printf("Part 2: %d\n", otherSum)
}

func StepAll(foundTrailends *map[vec.Vec]struct{}, vm *vec.VecMap, v vec.Vec) int {
	if vm.Vm[v] == '9' {
		(*foundTrailends)[v] = struct{}{}
		return 1
	}

	sum := 0
	for _, nv := range StepUp(vm, v) {
		sum += StepAll(foundTrailends, vm, nv)
	}
	return sum
}

func StepUp(vm *vec.VecMap, v vec.Vec) []vec.Vec {
	out := make([]vec.Vec, 0)
	h := vm.Vm[v]
	for _, o := range vm.NeighborsTo(v) {
		if vm.Vm[o] - h == 1 {
			out = append(out, o)
		}
	}
	return out
}

func part_2(input string) {
	part_1(input)
}
