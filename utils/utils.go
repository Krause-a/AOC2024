package utils

import (
	"fmt"
	"os"
	"strings"
)

func Test() {
	fmt.Println("This is a test!")
}

func ParseInput() {
	if len(os.Args) != 2 {
		fmt.Println("Expected a single argument 1 or 2 to represent which part of the day is to be ran")
		os.Exit(1)
	}
	part := strings.Trim(os.Args[1], " \n")
	if part != "1" && part != "2" {
		fmt.Println("Expected a single argument 1 or 2 to represent which part of the day is to be ran")
		os.Exit(1)
	}
	fmt.Println(part)
}
