package main

import (
	y2023 "advent-of-go/solutions/2023"
	"advent-of-go/utils"
	"fmt"
)

func main() {
	for _, s := range y2023.Solutions() {
		printSolution(s)
	}
}

func printSolution(s utils.Solution) {
	fmt.Printf("%s: %s\n", s.Name(), stringifyRes(s.Calculate()))
}

func stringifyRes(s string, e error) string {
	if e != nil {
		return fmt.Sprintf("Error: %+v", e)
	}
	return s
}