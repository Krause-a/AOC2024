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

type FreeSpace struct {
	start int
	stop int
}

func (fs FreeSpace) Width() int {
	return fs.stop - fs.start + 1
}

type FileSize struct {
	start int
	stop int
	size int
}

func (fs FileSize) Width() int {
	return fs.stop - fs.start + 1
}

func part_2(input string) {

	files := make([]FileSize, 0)
	freeSpaces := make([]FreeSpace, 0)

	currentIndex := 0

	for i, char := range input {
		size, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		if i & 1 == 0 { // File Space
			files = append(files, FileSize {
				start: currentIndex,
				stop: currentIndex + size - 1,
				size: i/2,
			})
		} else { // Free Space
			freeSpaces = append(freeSpaces, FreeSpace{
				start: currentIndex,
				stop: currentIndex + size - 1,
			})
			currentIndex += size
		}
	}

	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]
		foundSpaceIndex := -1
		var space FreeSpace
		for spaceIndex, freeSpace := range freeSpaces {
			if freeSpace.Width() >= file.Width() && freeSpace.start < file.start {
				foundSpaceIndex = spaceIndex
				space = freeSpace
				break
			}
		}
		if foundSpaceIndex >= 0 {
			file.stop = space.start + file.Width() - 1
			file.start = space.start
			space.start = file.stop + 1
			files[i] = file
			freeSpaces[foundSpaceIndex] = space
		}
	}

	fmt.Printf("%v\n", files)
	printFiles(files)

	// Got the fill order wrong here
	// for _, freeSpace := range freeSpaces {
	// 	screw_it:
	// 	rightIndex := len(files) - 1
	// 	fileToReindex := files[rightIndex]
	// 	println()
	// 	for true {
	// 		fileToReindex = files[rightIndex]
	// 		fmt.Printf("%v, %v\n", freeSpace, fileToReindex)
	// 		if fileToReindex.Width() <= freeSpace.Width() && fileToReindex.start > freeSpace.stop {
	// 			break
	// 		}
	// 		rightIndex -= 1
	// 		if rightIndex < 0 {
	// 			panic("huh?")
	// 		}
	// 	}
	// 	if fileToReindex.Width() > freeSpace.Width() || fileToReindex.start < freeSpace.stop {
	// 		fmt.Printf("Couldn't find one of interest\n")
	// 		break
	// 	}
	// 	fmt.Printf("Selected free: %v, file: %v\n", freeSpace, fileToReindex)
	// 	fmt.Printf("free width: %d, file width: %d\n", freeSpace.Width(), fileToReindex.Width())
	// 	// Order
	// 	fileToReindex.stop = freeSpace.start + fileToReindex.Width() - 1
	// 	fileToReindex.start = freeSpace.start
	// 	//
	// 	freeSpace.start = fileToReindex.stop + 1
	// 	fmt.Printf("free width: %d, file width: %d\n", freeSpace.Width(), fileToReindex.Width())
	// 	if freeSpace.Width() != 0 {
	// 		println("screw it")
	// 		goto screw_it
	// 	}
	// }

	sum := 0
	for _, file := range files {
		fmt.Printf("file: %v\nSmurt Maff = %d\n", file, ((file.Width() * (file.Width() - 1))/2 + file.start * file.Width()) * file.size)
		sum += ((file.Width() * (file.Width() - 1))/2 + file.start * file.Width()) * file.size

		dumbMaff := 0
		for i := file.start; i <= file.stop; i++ {
			dumbMaff += file.size * i
		}
		fmt.Printf("Dumb Maff = %d\n", dumbMaff)
	}

	println(sum)
	//   89536572931 is too low
	//     853653744 is too low dummy
	// 6337367222422

}

func printFiles(files []FileSize) {
	// This has shown a bug!
	// For some reason I am having duplicate start and stop indicies for sizes 1 and 7.
	index := 0
	for true {
		flag := false
		largest := 0
		for _, file := range files {
			largest = max(largest, file.stop)
			if file.start <= index && index <= file.stop {
				fmt.Printf("%d", file.size)
				flag = true
				break
			}
		}
		if !flag && largest < index{
			break
		}
		if !flag {
			fmt.Print(".")
		}
		index += 1
	}
}
