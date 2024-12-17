package main

import (
	"aoc_2024/utils"
	"aoc_2024/utils/datastructures"
	"fmt"
	"math"
	"strconv"
	"strings"
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

type Stone int

func part_1(input string) {
	TheWholeThing(input, 25)
}

func (s Stone) FirstRule() (bool, Stone) {
	if s == 0 {
		// fmt.Printf("First Rule Hit with %d\n", s)
		return true, Stone(1)
	} else {
		return false, 0
	}
}

func (s Stone) SecondRule() (bool, [2]Stone) {
	digits := 0
	divRunner := s
	for divRunner > 0 {
		digits += 1
		divRunner = divRunner / 10
	}
	if digits & 1 == 0 {
		left := int(s) / int(math.Pow10(digits / 2))
		right := int(s) % int(math.Pow10(digits / 2))
		// fmt.Printf("Second Rule Hit with %d    %d%d\n", s, left, right)
		return true, [2]Stone{Stone(left), Stone(right)}
	} else {
		return false, [2]Stone{}
	}
}

func (s Stone) ThirdRule() Stone {
	// fmt.Printf("Third Rule Hit with %d\n", s)
	return s * 2024
}

func part_2(input string) {
	TheWholeThing(input, 75)
}

func TheWholeThing(input string, blinks int) {
	numberStrings := strings.Split(input, " ")
	stones := make([]Stone, 0)
	for _, numberStr := range numberStrings {
		stone, err := strconv.Atoi(numberStr)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("Adding stone %d\n", stone)
		stones = append(stones, Stone(stone))
	}
	list := datastructures.MakeLinkedList(stones)
	// list.PrintList()
	for range blinks {
		tail := list.Head
		// fmt.Printf("TAIL\n")
		consumed, newStone := tail.Value.FirstRule()
		if consumed {
			tail.Value = newStone
		} else {
			consumed, newStones := tail.Value.SecondRule()
			if consumed {
				list.Head = &datastructures.Node[Stone] {
					Value: newStones[0],
					Next: tail,
				}
				list.Size += 1
				tail.Value = newStones[1]
			} else {
				tail.Value = tail.Value.ThirdRule()
			}
		}

		next := tail.Next
		for next != nil {
			// fmt.Printf("Testing!\n")
			consumed, newStone := next.Value.FirstRule()
			if consumed {
				next.Value = newStone
			} else {
				consumed, newStones := next.Value.SecondRule()
				if consumed {
					tail.Next = &datastructures.Node[Stone] {
						Value: newStones[0],
						Next: next,
					}
					list.Size += 1
					next.Value = newStones[1]
				} else {
					next.Value = next.Value.ThirdRule()
				}
			}
			tail = next
			next = next.Next
		}
		// fmt.Printf("LAST!\n")
		// list.PrintList()
	}


	fmt.Printf("%d\n", list.Size)
}
