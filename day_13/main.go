package main

import (
	"aoc_2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, part := utils.ParseInput(13)
	if part == 1 {
		part_1(input)
	} else {
		part_2(input)
	}
}

type Point struct {
	x int
	y int
}

func (p *Point) add(o Point) {
	p.x += o.x
	p.y += o.y
}

func (p *Point) sub(o Point) {
	p.x -= o.x
	p.y -= o.y
}

func (p *Point) equals(o Point) bool {
	return p.x == o.x && p.y == o.y
}

func (p *Point) isPast(o Point) bool {
	return p.x > o.x || p.y > o.y
}

func (p *Point) toString() string {
	return fmt.Sprintf("{x: %d, y: %d}", p.x, p.y)
}

func squareDistance(p1 Point, p2 Point) int {
	var my_p = Point{x: p1.x, y: p1.y}
	my_p.sub(p2)
	return my_p.x * my_p.x + my_p.y * my_p.y
}

type Button = Point
type Prize = Point
type Token = int

type Machine struct {
	a Button
	b Button
	prize Prize
	current_location Point
	a_tokens Token
	b_tokens Token
	got_prize bool
}

func makeMachine(a Button, b Button, prize Prize) *Machine {
	return &Machine {
		a: a,
		b: b,
		prize: prize,
		current_location: Point{x: 0, y: 0},
		a_tokens: 0,
		b_tokens: 0,
		got_prize: false,
	}
}

func makeMachineFromLines(lines [3]string) *Machine {
	var r = regexp.MustCompile("[0-9]+")
	var a_parts = r.FindAllString(lines[0], 2)
	var b_parts = r.FindAllString(lines[1], 2)
	var prize_parts = r.FindAllString(lines[2], 2)
	var a_x, _ = strconv.Atoi(a_parts[0])
	var a_y, _ = strconv.Atoi(a_parts[1])
	var b_x, _ = strconv.Atoi(b_parts[0])
	var b_y, _ = strconv.Atoi(b_parts[1])
	var prize_x, _ = strconv.Atoi(prize_parts[0])
	var prize_y, _ = strconv.Atoi(prize_parts[1])

	var a = Button{x: a_x, y: a_y}
	var b = Button{x: b_x, y: b_y}
	var prize = Prize{x: prize_x, y: prize_y}
	return makeMachine(a, b, prize)
}

type MachineAction int

const (
	MachineActionA = iota
	MachineActionB
)

func (m *Machine) spent() Token {
	return m.a_tokens + m.b_tokens
}

func (m *Machine) isPastPrize() bool {
	return m.current_location.isPast(m.prize)
}

func (m *Machine) isOnPrize() bool {
	if m.current_location.equals(m.prize) {
		m.got_prize = true
		return true
	}
	return false
}

func (m *Machine) pressA() {
	m.a_tokens += 3
	m.current_location.add(m.a)
}

func (m *Machine) unpressA() {
	m.a_tokens -= 3
	m.current_location.sub(m.a)
}

func (m *Machine) pressB() {
	m.b_tokens += 1
	m.current_location.add(m.b)
}

func (m *Machine) unpressB() {
	m.b_tokens -= 1
	m.current_location.sub(m.b)
}

// Idea: 1. Press the cheaper button until at or past the prize. 2. Subtract the cheaper button until before the prize. 3. Add the expensize button until at or past the prize. 4. Repeat from 2 until at the prize.

func part_1(input string) {
	println("Part 1 START")
	var machines = lines_into_machines(input)
	println(len(machines))
	var saftey = 1_000_000
	for i, machine := range machines {
		var machine_saftey = 1_000
		fmt.Printf("In machine %d\n", i)
		for !machine.isPastPrize() {
			machine.pressB()
		}
		fmt.Printf("Spent %d tokens on B for %d\n", machine.spent, i)
		past_by_x := machine.current_location.x > machine.prize.x
		past_by_y := machine.current_location.y > machine.prize.y
		for saftey > 0 {
			saftey -= 1
			machine_saftey -= 1
			if machine.isOnPrize() {
				fmt.Printf("Machine %d found the prize in %d tokens\n", i, machine.spent)
				break
			}
			for machine.isPastPrize() {
				machine.unpressB()
			}
			if machine.isOnPrize() {
				fmt.Printf("Machine %d found the prize in %d tokens\n", i, machine.spent)
				break
			}
			machine.pressA()
			if past_by_x && machine.current_location.y > machine.prize.y || past_by_y && machine.current_location.x > machine.prize.x || machine.b_tokens < 0 {
				if machine_saftey < 1 {
					fmt.Printf("Machine %d has no solution\n", i)
					break
				}
			}
		}
		if saftey == 0 {
			fmt.Printf("%v\n", machine)
			println("\n\n\n\033[1;31mSAFTEY TRIPPED!!!\033[0m\n\n\n")
			return
		} else {
			fmt.Printf("Machine Done %d, Saftey remaining: %d\n", i, saftey)
		}
	}
	var total_tokens Token = 0
	for i, machine := range machines {
		if machine.got_prize {
			fmt.Printf("Spent %d tokens to get prize for %d\n", machine.spent(), i )
			total_tokens += machine.spent()
		}
	}
	fmt.Printf("Total token spent to all prizes: %d\n", total_tokens)
	println("Part 1 END")
}
// ANSWER ATTEMP: 32703 (Too low)
// ANSWER ATTEMP: 31519 (Too low)
// ANSWER ATTEMP: 32163 (Too low)
// ANSWER ATTEMP: 35574 (Too low)

func lines_into_machines(input string) []*Machine {
	lines := strings.Split(input, "\n")
	var machines = make([]*Machine, 0)
	var machine_lines [3]string
	var machine_index = 0
	for _, line := range lines {
		if machine_index == 3 {
			machines = append(machines, makeMachineFromLines(machine_lines))
			machine_index = 0
		} else {
			machine_lines[machine_index] = line
			machine_index += 1
		}
	}
	if machine_index == 3 {
		machines = append(machines, makeMachineFromLines(machine_lines))
		machine_index = 0
	}
	return machines
}

func part_2(input string) {
	println("Part 2 START")
	var machines = lines_into_machines(input)
	println(len(machines))
	var saftey = 1_000_000
	var prize_offset = Point{x: 10000000000000, y: 10000000000000}
	// Obv this will timeout.
	for i, machine := range machines {
		machine.prize.add(prize_offset)
		var machine_saftey = 1_000
		fmt.Printf("In machine %d\n", i)
		for !machine.isPastPrize() {
			machine.pressB()
		}
		fmt.Printf("Spent %d tokens on B for %d\n", machine.spent, i)
		past_by_x := machine.current_location.x > machine.prize.x
		past_by_y := machine.current_location.y > machine.prize.y
		for saftey > 0 {
			saftey -= 1
			machine_saftey -= 1
			if machine.isOnPrize() {
				fmt.Printf("Machine %d found the prize in %d tokens\n", i, machine.spent)
				break
			}
			for machine.isPastPrize() {
				machine.unpressB()
			}
			if machine.isOnPrize() {
				fmt.Printf("Machine %d found the prize in %d tokens\n", i, machine.spent)
				break
			}
			machine.pressA()
			if past_by_x && machine.current_location.y > machine.prize.y || past_by_y && machine.current_location.x > machine.prize.x || machine.b_tokens < 0 {
				if machine_saftey < 1 {
					fmt.Printf("Machine %d has no solution\n", i)
					break
				}
			}
		}
		if saftey == 0 {
			fmt.Printf("%v\n", machine)
			println("\n\n\n\033[1;31mSAFTEY TRIPPED!!!\033[0m\n\n\n")
			return
		} else {
			fmt.Printf("Machine Done %d, Saftey remaining: %d\n", i, saftey)
		}
	}
	var total_tokens Token = 0
	for i, machine := range machines {
		if machine.got_prize {
			fmt.Printf("Spent %d tokens to get prize for %d\n", machine.spent(), i )
			total_tokens += machine.spent()
		}
	}
	fmt.Printf("Total token spent to all prizes: %d\n", total_tokens)
	println("Part 2 END")
}

