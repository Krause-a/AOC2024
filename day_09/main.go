package main

import (
	"aoc_2024/utils"
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

type File struct {
	id int
	size int
}
type FileSlice []File

func (fs FileSlice) Len() int {
	return len(fs)
}

func (fs FileSlice) Less(i, j int) bool {
	return fs[i].id < fs[j].id
}

func (fs FileSlice) Swap(i, j int) {
	fs[i], fs[j] = fs[j], fs[i]
}

func part_1(input string) {

	files := make(FileSlice, 0)
	freeIndicies := make([]int, 0)

	currentIndex := 0

	for i, char := range input {
		size, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		if i & 1 == 0 { // File Space
			for range size {
				files = append(files, File{
					id: currentIndex,
					size: i/2,
				})
				currentIndex += 1
			}
		} else { // Free Space
			for range size {
				freeIndicies = append(freeIndicies, currentIndex)
				currentIndex += 1
			}
		}
	}

	rightIndex := len(files) - 1
	for _, freeIndex := range freeIndicies {
		if files[rightIndex].id <= freeIndex {
			break
		}
		files[rightIndex].id = freeIndex
		rightIndex -= 1
	}

	sum := 0
	for _, file := range files {
		// fmt.Printf("%d * %d = %d\n", file.id, file.size, int(file.id) * int(file.size))
		sum += int(file.id) * int(file.size)
	}

	println(sum)
	//   89536572931 is too low
	//     853653744 is too low dummy
	// 6337367222422
}

func part_2(input string) {
	println("Part 2 START")
	println(input)
	println("Part 2 END")
}
