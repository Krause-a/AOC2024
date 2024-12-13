package main

import (
	"aoc_2024/utils"
	"fmt"
	"strings"
)

func main() {
	fmt.Print("")
	input, part := utils.ParseInput(2)
	if part == 1 {
		part_1(input)
	} else {
		part_2(input)
	}
}

func part_1(input string) {
	size, blocks, gPos, gDir := parseInputMap(input)
	guard := makeGuard(size, &blocks, gPos, gDir)
	// var stepCount int
	for !guard.isRepeatState() && guard.isInArea() {
		// stepCount++
		// fmt.Printf("\nStep %d\n", stepCount)
		// guard.printInfo()
		// if stepCount >= 10 {
		// 	return
		// }
		// fmt.Printf("current guard position %v\n", guard.pos)
		guard.takeStep()
	}
	println(len(guard.history))
}

type BlockMap map[Vec]struct{}
type StepHistory map[Vec][4]bool

type Vec struct {
	x int
	y int
}

type Guard struct {
	blockMap *BlockMap
	history StepHistory
	mapSize Vec
	pos Vec
	dir Vec
}

func makeGuard(mapSize Vec, blockMap *BlockMap, pos Vec, dir Vec) Guard {
	history := make(StepHistory, 0)
	return Guard {
		pos: pos,
		mapSize: mapSize,
		blockMap: blockMap,
		dir: dir,
		history: history,
	}
}

func (g *Guard) takeStep() {
	directions, exists := g.history[g.pos]
	if exists {
		directions[dirToIndex(g.dir)] = true
		g.history[g.pos] = directions
	} else {
		var dirs [4]bool
		dirs[dirToIndex(g.dir)] = true
		g.history[g.pos] = dirs
	}
	for g.checkBlocked() {
		// fmt.Printf("BLOCKED:")
		g.dir = rotate(g.dir)
	}
	// fmt.Printf("Previous %v -> ", g.pos)
	g.pos = addVec(g.pos, g.dir)
	// fmt.Printf("%v\n", g.pos)
	// fmt.Printf("Step Taken %v\n", g.pos)
}
func (g *Guard) wouldBeNextPos() Vec {

	dir := g.dir
	nextPos := addVec(g.pos, dir)
	_, blocked := (*g.blockMap)[nextPos]
	for blocked {
		// fmt.Printf("BLOCKED:")
		dir = rotate(dir)
		nextPos = addVec(g.pos, dir)
		_, blocked = (*g.blockMap)[nextPos]
	}
	return nextPos
}

func (g *Guard) checkBlocked() bool {
	// fmt.Printf("Check %v + %v\n", g.pos, g.dir)
	nextPos := addVec(g.pos, g.dir)
	_, blocked := (*g.blockMap)[nextPos]
	// fmt.Printf("Is blocked? %v\n", blocked)
	return blocked
}
func (g *Guard) isInArea() bool {
	return g.pos.x >= 0 && g.pos.y >= 0 && g.pos.x < g.mapSize.x && g.pos.y < g.mapSize.y
}
func (g *Guard) isRepeatState() bool {
	dirs, exists := g.history[g.pos]
	// fmt.Printf("history %v check %v\n", g.history, g.pos)
	if exists {
		// fmt.Printf("State exists %v[%d]\n", dirs, dirToIndex(g.dir))
		return dirs[dirToIndex(g.dir)]
	}
	return false
}
func (g *Guard) printInfo() {
	var output string
	for y := range g.mapSize.y {
		for x := range g.mapSize.x {
			pos := Vec {x: x, y: y}
			if pos == g.pos {
				var dirChar string
				if g.dir == Up {
					dirChar = "^ "
				} else if g.dir == Down {
					dirChar = "v "
				} else if g.dir == Left {
					dirChar = "< "
				} else if g.dir == Right {
					dirChar = "> "
				}
				output += dirChar
				continue
			}
			_, blocked := (*g.blockMap)[pos]
			if blocked {
				output += "# "
				continue
			}
			dirs, stepped := g.history[pos]
			if stepped {
				var dirStr string
				if dirs[dirToIndex(Up)] {
					dirStr = "^ "
				} else if dirs[dirToIndex(Down)] {
					dirStr = "v "
				}
				if dirs[dirToIndex(Left)] {
					if len(dirStr) > 0 {
						dirStr = "+ "
					} else {
						dirStr = "< "
					}
				} else if dirs[dirToIndex(Right)] {
					if len(dirStr) > 0 {
						dirStr = "+ "
					} else {
						dirStr = "> "
					}
				}
				output += dirStr
				continue
			}
			output += ". "
		}
		output += "\n"
	}
	fmt.Print(output)
}

var Up = Vec{x: 0, y: -1}
var Down = Vec{x: 0, y: 1}
var Left = Vec{x: -1, y: 0}
var Right = Vec{x: 1, y: 0}

func dirToIndex(v Vec) int {
	if v == Up {
		return 0
	} else if v == Down {
		return 1
	} else if v == Left {
		return 2
	} else if v == Right {
		return 3
	}
	panic("None const direction attempted conversion to index")
}

func addVec(v1 Vec, v2 Vec) Vec {
	return Vec {
		x: v1.x + v2.x,
		y: v1.y + v2.y,
	}
}

func rotate(v Vec) Vec {
	return Vec {
		x: -v.y,
		y: v.x,
	}
}

func invertVec(v Vec) Vec {
	return Vec {
		x: -v.x,
		y: -v.y,
	}
}

// Max vector, location of obstructions, guard position, guard direction
func parseInputMap(input string) (Vec, BlockMap, Vec, Vec) {
	var pos Vec
	var dir Vec
	blocks := make(map[Vec]struct{}, 0)
	var max Vec
	for y, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		max.y = y + 1
		max.x = len(line)
		for x := range len(line) {
			if line[x] == '#' {
				blocks[Vec{x: x, y: y}] = struct{}{}
			} else if strings.Contains("^v<>", string(line[x])) {
				pos = Vec{
					x: x,
					y: y,
				}
				dir = parseDirection(line[x])
			}
		}
	}
	return max, blocks, pos, dir
}

func parseDirection(char byte) Vec {
	switch char {
	case '^':
		return Up
	case 'v':
		return Down
	case '<':
		return Left
	case '>':
		return Right
	}
	panic("Given an invalid direction")
}


func part_2(input string) {
	clints_way := true
	if clints_way {
		clint_way(input)
		return
	}
	size, blocks, gPos, gDir := parseInputMap(input)
	loopBlocks := make(BlockMap)
	guard := makeGuard(size, &blocks, gPos, gDir)
	for !guard.isRepeatState() && guard.isInArea() {
		nextPos := addVec(guard.pos, guard.dir)
		dir := guard.dir
		_, blocked := blocks[nextPos]
		for blocked {
			dir = rotate(dir)
			nextPos = addVec(guard.pos, dir)
			_, blocked = blocks[nextPos]
		}
		_, stepped := guard.history[nextPos]
		if !blocked && !stepped && guard.willLoopIfBlocked() {
		// if guard.willLoopIfBlocked() {
			loopBlocks[nextPos] = struct{}{}
		}
		guard.takeStep()
	}
	guard.printInfo()
	println()

	var printStr string
	for y := range size.y {
		for x := range size.x {
			pos := Vec {x: x, y: y}
			_, exists := loopBlocks[pos]
			if exists {
				printStr += "O "
			} else {
				_, exists = blocks[pos]
				if exists {
					printStr += "# "
				} else {
					printStr += ". "
				}
			}
		}
		printStr += "\n"
	}

	println(printStr)

	println(len(loopBlocks))
	// 1883 is too high
	// 500 is too low
	// Clint says somewhere around over 1500
	// 1785 is too low
	// 1783 is too low
	// 1792 is wrong (no longer low or high)
}


func (g *Guard) willLoopIfBlocked() bool {
	hypotheticalBlockPos := addVec(g.dir, g.pos)
	// guard starts out rotated
	ghostGuard := makeGuard(g.mapSize, g.blockMap, g.pos, g.dir)
	for !ghostGuard.isRepeatState() && ghostGuard.isInArea() {
		nextStep := addVec(ghostGuard.dir, ghostGuard.pos)
		if nextStep == hypotheticalBlockPos {
			ghostGuard.dir = rotate(ghostGuard.dir)
		}
		ghostGuard.takeStep()
		if ghostGuard.pos == hypotheticalBlockPos {
			return true // stuck in 4 blocks
		}
		// ghostGuard.printInfo()
	}
	// if ghostGuard.isRepeatState() {
	// 	println("Ghost Path")
	// 	ghostGuard.printInfo()
	// 	println()
	// }
	return ghostGuard.isRepeatState()
}


func clint_way(input string) {
	fmt.Printf("Clint's Way\n")
	size, blocks, gPos, gDir := parseInputMap(input)
	loopBlocks := make(BlockMap)
	guard := makeGuard(size, &blocks, gPos, gDir)
	for !guard.isRepeatState() && guard.isInArea() {
		guard.takeStep()
	}
	// guard.printInfo()
	// println()

	for pos, dirs := range guard.history {
		for _, isTrue:= range dirs {
			if !isTrue {
				continue
			}
			// ghostGuardPos := addVec(pos, invertVec(indexToDir(dir)))

			hypoBlockPos := pos

			ghostGuard := makeGuard(size, &blocks, gPos, gDir)
			for !ghostGuard.isRepeatState() && ghostGuard.isInArea() {

				for i := 0; ghostGuard.wouldBeNextPos() == hypoBlockPos && i < 4; i++ {
					ghostGuard.dir = rotate(ghostGuard.dir)
					// fmt.Printf("Hypo deflect\n")
				}
				ghostGuard.takeStep()
				if ghostGuard.pos == hypoBlockPos {
					// stuck in 4 blocks
					fmt.Printf("4 blocks with %v\n", hypoBlockPos)
					loopBlocks[hypoBlockPos] = struct{}{}
					// ghostGuard.printInfo()
					// println()
					break
				}
				// ghostGuard.printInfo()
			}
			if ghostGuard.isRepeatState() {
				loopBlocks[hypoBlockPos] = struct{}{}
			}
		}
	}

	var printStr string
	for y := range size.y {
		for x := range size.x {
			pos := Vec {x: x, y: y}
			_, exists := loopBlocks[pos]
			if exists {
				printStr += "O "

			} else if pos == gPos {
				if gDir == Up {
					printStr += "^ "
				} else if gDir == Down {
					printStr += "v "
				} else if gDir == Left {
					printStr += "< "
				} else if gDir == Right {
					printStr += "> "
				}
			}else {
				_, exists = blocks[pos]
				if exists {
					printStr += "# "
				} else {
					printStr += ". "
				}
			}
		}
		printStr += "\n"
	}

	println(printStr)

	println(len(loopBlocks))
	// 1883 is too high
	// 500 is too low
	// Clint says somewhere around over 1500
	// 1785 is too low
	// 1783 is too low
	// 1792 is wrong (no longer low or high)
}

func indexToDir(i int) Vec {
	if i == 0 {
		return Up
	} else if i == 1 {
		return Down
	} else if i == 2 {
		return Left
	} else if i == 3 {
		return Right
	}
	panic("None const direction attempted conversion to index")
}
