package main

import (
	"aoc_2024/utils"
	"aoc_2024/utils/vec"
	"fmt"
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
	freqs, bounds := parseInput(input)
	antinodes := make(map[vec.Vec]struct{})
	for _, vecs := range freqs {
		lines := make([]vec.Line, 0)
		for i, v := range vecs {
			a := v
			for j := range (len(vecs) - i) {
				b := vecs[j + i]
				if a == b {
					continue
				}
				lines = append(lines, vec.MakeLine(a, b))
			}
		}

		for _, line := range lines {
			fmt.Printf("%v, %v => %v\n", line.A, line.B, line.AToB())
			aToB := line.AToB()
			bToA := line.BToA()
			antiPoint := aToB.Add(line.B)
			if bounds.Contains(antiPoint) {
				antinodes[antiPoint] = struct{}{}
			}
			antiPoint = bToA.Add(line.A)
			if bounds.Contains(antiPoint) {
				antinodes[antiPoint] = struct{}{}
			}
		}
	}

	println(len(antinodes))
}

type Freq rune
type FreqMap map[Freq][]vec.Vec
func parseInput(input string) (FreqMap, vec.Box) {
	antenne := make(FreqMap)
	bounds := utils.EachRuneWithVec(input, func(v vec.Vec, r rune) {
		if r == '.' {
			return
		}
		antenne[Freq(r)] = append(antenne[Freq(r)], v)
	})
	fmt.Printf("%v\n", bounds)
	return antenne, vec.MakeBox(vec.Zero, bounds.Sub(vec.One))
}

func part_2(input string) {
	freqs, bounds := parseInput(input)
	antinodes := make(map[vec.Vec]struct{})
	for _, vecs := range freqs {
		lines := make([]vec.Line, 0)
		for i, v := range vecs {
			a := v
			for j := range (len(vecs) - i) {
				b := vecs[j + i]
				if a == b {
					continue
				}
				lines = append(lines, vec.MakeLine(a, b))
			}
		}

		for _, line := range lines {
			// fmt.Printf("%v ===> %v\n", line.A, line.B)
			aToBProgress := line.A
			aToB := line.AToB()
			aToBStep := aToB.NormalizeToInt()
			// fmt.Printf("%v -> %v\n", aToB, aToBStep)
			// fmt.Printf("%v", aToBProgress)
			for true {
				aToBProgress = aToBProgress.Add(aToBStep)
				if !bounds.Contains(aToBProgress) {
					break
				}
				antinodes[aToBProgress] = struct{}{}
				// fmt.Printf(" -> %v", aToBProgress)
			}
			// fmt.Printf("\n")
			bToAProgress := line.B
			bToA := line.BToA()
			bToAStep := bToA.NormalizeToInt()
			// fmt.Printf("%v -> %v\n", bToA, bToAStep)
			for true {
				bToAProgress = bToAProgress.Add(bToAStep)
				if !bounds.Contains(bToAProgress) {
					break
				}
				antinodes[bToAProgress] = struct{}{}
			}
		}
	}

	println(len(antinodes))
}
