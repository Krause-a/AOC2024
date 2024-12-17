package vec

import (
	"fmt"
	"strings"
)

type VecMap struct {
	Vm map[Vec]rune
	size Vec
}

func ParseIntoMap(input string) VecMap {
	lines := strings.Split(input, "\n")
	vm := VecMap {
		Vm: make(map[Vec]rune),
		size: Vec{x: len(lines[0]), y: len(lines)},
	}
	for y, line := range lines {
		for x, char := range line {
			vm.Vm[Vec{x: x, y: y}] = char
		}
	}
	return vm
}

func (vm *VecMap) NeighborsTo(v Vec) []Vec {
	n := make([]Vec, 0)
	for _, neighbor := range v.Neighbors() {
		if neighbor.x < 0 || neighbor.y < 0 || neighbor.x >= vm.size.x  || neighbor.y >= vm.size.y {
			continue
		}
		n = append(n, neighbor)
	}
	return n
}

func (vm *VecMap) Print() {
	for y := range vm.size.y {
		for x := range vm.size.x {
			fmt.Printf("%c ", vm.Vm[MakeVec(x, y)])
		}
		fmt.Println()
	}
}
