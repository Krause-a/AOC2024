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
	// TheWholeThing(input, 25) // 55312
	DynamicMethod(input, 25)
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
	// TheWholeThing(input, 75)
	DynamicMethod(input, 75)
}

type NextStone struct {
	ns1 Stone
	has2 bool
	ns2 Stone
}

type AgedStone struct {
	s Stone
	age int
}

func DynamicMethod(input string, blinks int) {
	// Well this did nothing notable and I ran out of RAM
	numberStrings := strings.Split(input, " ")
	startingStones := make([]AgedStone, 0)
	for _, numberStr := range numberStrings {
		stone, err := strconv.Atoi(numberStr)
		if err != nil {
			panic(err)
		}
		startingStones = append(startingStones, AgedStone {
			s: Stone(stone),
			age: 0,
		})
	}

	stoneHistory := make(map[Stone]NextStone)

	index := 0
	for index < len(startingStones) {
		as := startingStones[index]
		index += 1
		for as.age < blinks {
			consumed, ns := as.s.FirstRule()
			as.age += 1
			if consumed {
				stoneHistory[as.s] = NextStone{
					ns1: ns,
				}
			} else {
				consumed, nns := as.s.SecondRule()
				if consumed {
					stoneHistory[as.s] = NextStone{
						ns1: nns[0],
						has2: true,
						ns2: nns[1],
					}
					startingStones = append(startingStones, AgedStone{
						s: nns[1],
						age: as.age,
					})
				} else {
					stoneHistory[as.s] = NextStone{
						ns1: as.s.ThirdRule(),
					}
				}
			}
			as.s = stoneHistory[as.s].ns1
		}
		if index & 1048575 == 1048575 {
			fmt.Printf("%d out of %d\n", index, len(startingStones))
		}
	}
	fmt.Printf("%v\n", len(startingStones))
}

func TheWholeThing(input string, blinks int) {
	// The fact that the puzzle says it is ordered at all is a red herring
	// Each stone is completely alone
	numberStrings := strings.Split(input, " ")
	previousStones := make([]Stone, 0)
	for _, numberStr := range numberStrings {
		stone, err := strconv.Atoi(numberStr)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("Adding stone %d\n", stone)
		previousStones = append(previousStones, Stone(stone))
	}

	nextStones := make([]Stone, 0)

	// fmt.Printf("%v\n", previousStones)
	for blink := range blinks {
		fmt.Printf("Blink %d Current size = %d\n", blink, len(previousStones))
		index := 0
		for _, s := range previousStones {
			consumed, ns := s.FirstRule()
			if len(nextStones) <= index {
				nextStones = append(nextStones, Stone(0))
			}
			if consumed {
				nextStones[index] = ns
			} else {
				consumed, nns := s.SecondRule()
				if consumed {
					nextStones[index] = nns[0]
					index += 1
					if len(nextStones) <= index {
						nextStones = append(nextStones, Stone(0))
					}
					nextStones[index] = nns[1]
				} else {
					nextStones[index] = s.ThirdRule()
				}
			}
			index += 1
		}
		nextStones, previousStones = previousStones, nextStones
		// fmt.Printf("%v\n", previousStones)
	}

	fmt.Printf("%d\n", len(previousStones))
}

func part_1_for_posterity(input string) {
	// The fact that the puzzle says it is ordered at all is a red herring
	// Each stone is completely alone
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
	for range 25 {
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
