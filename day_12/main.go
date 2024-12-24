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

type Region struct {
	plants []vec.Vec
	plantType rune
}

func makeRegion(pt rune) *Region {
	return &Region{
		plants: make([]vec.Vec, 0),
		plantType: pt,
	}
}

func part_1(input string) {
	fmt.Println()
	vm := vec.ParseIntoMap(input)
	regions := make([]*Region, 0)

	for v, p := range vm.Vm {
		// fmt.Printf("Make region\n")
		region := makeRegion(p)
		for _, r := range regions {
			if r.plantType == p && pointInMap(v, r) {
				region = nil
				// fmt.Printf("Found existing region\n")
				break
			}
		}
		if region == nil {
			continue
		}

		currentPoints := make([]vec.Vec, 0)
		nextPoints := make([]vec.Vec, 0)
		currentPoints = append(currentPoints, v)
		region.plants = append(region.plants, v)
		// fmt.Printf("Start region building\n")

		for len(currentPoints) > 0 {
			// fmt.Printf("Current Point GT 0 Loop\n")
			for _, cp := range currentPoints {
				// fmt.Printf("Current Point range Loop\n")
				for _, np := range vm.NeighborsTo(cp) {
					// fmt.Printf("Next Point range Loop\nCompare %v vs %v\n", cp, np)
					if vm.Vm[np] == region.plantType && !pointInMap(np, region) && !inSlice(np, nextPoints) {
						// fmt.Printf("Add point to region %v\n", np)
						nextPoints = append(nextPoints, np)
						region.plants = append(region.plants, np)
					}
				}
			}
			// clear(currentPoints)
			currentPoints = make([]vec.Vec, 0)
			nextPoints, currentPoints = currentPoints, nextPoints
			// fmt.Printf("%v bbbb %v\n\n", currentPoints, nextPoints)
		}
		regions = append(regions, region)
	}

	totalCost := 0
	for _, r := range regions {
		perimeter := 0
		area := len(r.plants)
		for _, point := range r.plants {
			missing := 4
			for _, np := range point.Neighbors() {
				if pointInMap(np, r) {
					missing -= 1
				}
			}
			perimeter += missing
		}
		cost := area * perimeter
		totalCost += cost
	}

	fmt.Printf("%d\n", totalCost)
}

func inSlice(v vec.Vec, s []vec.Vec) bool {
	for _, asdf := range s {
		if asdf == v {
			return true
		}
	}
	return false
}

func pointInMap(point vec.Vec, r *Region) bool {
	for _, ov := range r.plants {
		if ov == point {
			return true
		}
	}
	return false
}


func part_2(input string) {
	fmt.Println()
	vm := vec.ParseIntoMap(input)
	regions := make([]*Region, 0)

	for v, p := range vm.Vm {
		// fmt.Printf("Make region\n")
		region := makeRegion(p)
		for _, r := range regions {
			if r.plantType == p && pointInMap(v, r) {
				region = nil
				// fmt.Printf("Found existing region\n")
				break
			}
		}
		if region == nil {
			continue
		}

		currentPoints := make([]vec.Vec, 0)
		nextPoints := make([]vec.Vec, 0)
		currentPoints = append(currentPoints, v)
		region.plants = append(region.plants, v)
		// fmt.Printf("Start region building\n")

		for len(currentPoints) > 0 {
			// fmt.Printf("Current Point GT 0 Loop\n")
			for _, cp := range currentPoints {
				// fmt.Printf("Current Point range Loop\n")
				for _, np := range vm.NeighborsTo(cp) {
					// fmt.Printf("Next Point range Loop\nCompare %v vs %v\n", cp, np)
					if vm.Vm[np] == region.plantType && !pointInMap(np, region) && !inSlice(np, nextPoints) {
						// fmt.Printf("Add point to region %v\n", np)
						nextPoints = append(nextPoints, np)
						region.plants = append(region.plants, np)
					}
				}
			}
			// clear(currentPoints)
			currentPoints = make([]vec.Vec, 0)
			nextPoints, currentPoints = currentPoints, nextPoints
			// fmt.Printf("%v bbbb %v\n\n", currentPoints, nextPoints)
		}
		regions = append(regions, region)
	}

	// This section needs a rewrite. The way I have it written has order issues and will generate inconsistent outputs for the E & X test input
	totalCost := 0
	for _, r := range regions {
		area := len(r.plants)
		perimeterPoints := make(map[vec.Vec]struct{}, 0)
		for _, point := range r.plants {
			for _, np := range point.Neighbors() {
				if !pointInMap(np, r) {
					perimeterPoints[np] = struct{}{}
				}
			}
		}
		fmt.Printf("%c -> %v\n", r.plantType, perimeterPoints)

		// Figure out sides
		sides := 0
		for perimeterPoint := range perimeterPoints {
			_, exists := perimeterPoints[perimeterPoint]
			if !exists {
				continue
			}
			delete(perimeterPoints, perimeterPoint)
			sides += 1
			fmt.Printf("Increment Side (1) %v\n", perimeterPoint)
			dirs := [4]vec.Vec {
				vec.Up,
				vec.Down,
				vec.Left,
				vec.Right,
			}


			inRegionCount := 0
			inPerCount := 0
			for _, goingNeighbor := range perimeterPoint.Neighbors() {
				if pointInMap(goingNeighbor, r) {
					inRegionCount += 1
				}
				_, inPer := perimeterPoints[goingNeighbor]
				if inPer {
					inPerCount += 1
				}
			}
			// fmt.Printf("\nRegion stuff %d %d\n", inRegionCount, inPerCount)
			if inPerCount == 0 && inRegionCount > 1 {
				fmt.Printf("Increment Side (2) %v\n", perimeterPoint)
				sides += inRegionCount - 1
			}


			for _, dir := range dirs {
				// Before I delete
				// A: Is the one I am about to delete touching 2 pieces of the region?
				// B: Also is it touching two pieces of the perimeter?
				// If the answer is A && !B then sides++

				going := perimeterPoint.Add(dir)
				if pointInMap(going, r) {
					continue
				}
				_, exists = perimeterPoints[going]
				for exists {
					inRegionCount := 0
					inPerCount := 0
					for _, goingNeighbor := range going.Neighbors() {
						if pointInMap(goingNeighbor, r) {
							inRegionCount += 1
						}
						_, inPer := perimeterPoints[goingNeighbor]
						if inPer {
							inPerCount += 1
						}
					}
					// fmt.Printf("\nRegion stuff %d %d\n", inRegionCount, inPerCount)
					if inPerCount == 0 && inRegionCount > 1 {
						fmt.Printf("Increment Side (3) %v by %d --- %v\n", going, inRegionCount - 1, dir)
						sides += inRegionCount - 1
					}
					
					delete(perimeterPoints, going)
					// fmt.Printf("   delete: %v", going)
					going = going.Add(dir)
					_, exists = perimeterPoints[going]
					if pointInMap(going, r) {
						break
					}
				}
			}
		}
		fmt.Printf("\nSides: %d\n\n", sides)

		cost := area * sides
		totalCost += cost
	}

	fmt.Printf("%d\n", totalCost)

}
