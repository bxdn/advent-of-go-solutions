package main

import (
	"advent-of-go/solutions"
	"advent-of-go/utils"
	"fmt"
)

func main() {
	for _, solutionSet := range solutions.Solutions() {
		for _, s := range solutionSet.Solutions {
			printSolution(s)
		}
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