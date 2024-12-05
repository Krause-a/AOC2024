package utils

import (
	"fmt"
	"os"
	"strings"
)

func ParseInput(day int) (string, int){
	input_err_string := "Expected a two arguments. t/r and 1/2 for test/run and part 1/part 2. Example `go run main.go t 2` for run test for part 2"
	if len(os.Args) != 3 {
		fmt.Println(input_err_string)
		os.Exit(1)
	}

	run_type := strings.Trim(os.Args[1], " \n")
	if run_type != "t" && run_type != "r" {
		fmt.Println(input_err_string)
		os.Exit(1)
	}

	part := strings.Trim(os.Args[2], " \n")
	if part != "1" && part != "2" {
		fmt.Println(input_err_string)
		os.Exit(1)
	}

	var part_num int
	if part == "1" {
		part_num = 1
	} else {
		part_num = 2
	}

	var input string
	if run_type == "t" {
		input = readTest(day, part_num)
	} else {
		input = readInput(day)
	}

	return input, part_num
}

func readTest(day int, part int) string {
	return readFile(day, fmt.Sprintf("test_%d", part))
}

func readInput(day int) string {
	return readFile(day, "input")
}

func readFile(day int, filename string) string {
	file_str, err_1 := os.ReadFile(fmt.Sprintf("day_%.2d/%s", day, filename))

	if err_1 != nil {
		file_str_2, err_2 := os.ReadFile(filename)
		if err_2 != nil {
			fmt.Println(err_2)
			os.Exit(3)
		}
		file_str = file_str_2
	}
	return string(file_str)
}

func Contains[T comparable](s []T, value T) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

func Any[T interface{}](s []T, fn func(T) bool) bool {
	for _, v := range s {
		if fn(v) {
			return true
		}
	}
	return false
}

func All[T interface{}](s []T, fn func(T) bool) bool {
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}
