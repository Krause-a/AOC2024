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
	tests := parseInput(input)
	passingTest := make([]TestInfo, 0)
	for _, test := range tests {
		ops := make([]OP, len(test.inputs) - 1)
		for i := range ops {
			ops[i] = ADD
		}
		for true {
			checkValue := 0
			for i, input := range test.inputs {
				if i == 0 {
					checkValue = input
				} else {
					checkValue = applyOperation(checkValue, input, ops[i-1])
				}
			}
			if checkValue == test.testValue {
				// fmt.Printf("Pass %d = %d ... %v . %v\n", test.testValue, checkValue, test.inputs, ops)
				passingTest = append(passingTest, test)
				break
			}

			hasAdd := false
			for _, op := range ops {
				if op == ADD {
					hasAdd = true
				}
			}
			if !hasAdd {
				break
			}
			for i, op := range ops {
				if op == ADD {
					ops[i] = MUL
					break
				} else if op == MUL {
					ops[i] = ADD
				}
			}
		}
	}

	sum := 0
	for _, test := range passingTest {
		sum += test.testValue
	}

	println(sum)
}

type OP rune
 const (
	ADD = '+'
	MUL = '*'
)

func applyOperation(a int, b int, op OP) int {
	if op == ADD {
		return a + b
	} else if op == MUL {
		return a * b
	} else {
		panic("Missing Op")
	}
}

type TestInfo struct {
	testValue int
	inputs []int
}

func parseTestInfoLine(line string) TestInfo {
	splits := strings.Split(line, ": ")
	testValue, err := strconv.Atoi(splits[0])
	if err != nil {
		panic(err)
	}
	inputsStrs := strings.Split(splits[1], " ")
	inputs := make([]int, 0)
	for _, str := range inputsStrs {
		val, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		inputs = append(inputs, val)
	}

	return TestInfo{
		testValue: testValue,
		inputs: inputs,
	}
}

func parseInput(input string) []TestInfo {
	lines := strings.Split(input, "\n")
	tests := make([]TestInfo, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		tests = append(tests, parseTestInfoLine(line))
	}
	return tests
}

func part_2(input string) {
	println("Part 2 START")
	println(input)
	println("Part 2 END")
}
