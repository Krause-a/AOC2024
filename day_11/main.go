package main

import (
	"aoc_2024/utils"
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
	// DynamicMethod(input, 25)
	ClintsWay(input, 25)
}

type StoneMap map[Stone]int

type Stonead struct {
	s1 Stone
	s2 Stone
	rule uint8
}

func (s Stone) IntoStonead() Stonead {
	return Stonead{
		s1: s,
	}
}

func (s Stonead) FirstRule() Stonead {
	if s.rule != 0 {
		return s
	}
	if s.s1 == 0 {
		// fmt.Printf("First Rule Hit with %d\n", s)
		return Stonead{
			s1: Stone(1),
			rule: 1,
		}
	} else {
		return s
	}
}

func (s Stonead) SecondRule() Stonead {
	if s.rule != 0 {
		return s
	}
	digits := 0
	divRunner := s.s1
	for divRunner > 0 {
		digits += 1
		divRunner = divRunner / 10
	}
	if digits & 1 == 0 {
		left := int(s.s1) / int(math.Pow10(digits / 2))
		right := int(s.s1) % int(math.Pow10(digits / 2))
		return Stonead{
			rule: 2,
			s1: Stone(left),
			s2: Stone(right),
		}
	} else {
		return s
	}
}

func (s Stonead) ThirdRule() Stonead {
	if s.rule != 0 {
		return s
	}
	return Stonead{
		rule: 3,
		s1: s.s1 * 2024,
	}
}

func (sm *StoneMap) InsertStoneadCount(s Stonead, c int) {
	prevCount := (*sm)[s.s1]
	(*sm)[s.s1] = prevCount + c
	if s.rule == 2 {
		prevCount = (*sm)[s.s2]
		(*sm)[s.s2] = prevCount + c
	}
}

func part_2(input string) {
	// TheWholeThing(input, 75)
	// DynamicMethod(input, 75)
	ClintsWay(input, 75)
}

func ClintsWay(input string, blinks int) {
	numberStrings := strings.Split(input, " ")
	stones := make(StoneMap)
	nextStones := make(StoneMap)
	for _, numberStr := range numberStrings {
		stone, err := strconv.Atoi(numberStr)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("Adding stone %d\n", stone)
		stones[Stone(stone)] = 1
	}

	// list.PrintList()
	for range blinks {
		for s, c := range stones {
			if c == 0 {
				continue
			}
			sad := s.IntoStonead().FirstRule().SecondRule().ThirdRule()
			nextStones.InsertStoneadCount(sad, c)
		}
		clear(stones)
		nextStones, stones = stones, nextStones
	}

	size := 0
	for _, c := range stones {
		size += c
	}

	fmt.Printf("Size = %d\n", size)
}
