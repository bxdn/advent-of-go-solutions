package main

import (
	"advent-of-go/solutions"
	"advent-of-go/utils"
	"flag"
	"fmt"
)

func main() {
	y := flag.Int("y", -1, "Year of solutions to display")
	d := flag.Int("d", -1, "Day of solutions to display")
	p := flag.Int("p", -1, "Part of solutions to display")
	t := flag.Bool("t", false, "Use to only test against known answers")
	q := flag.Bool("q", false, "Use to test in queit mode (only failures logged)")

	flag.Parse()

	solutionsToPrint := solutions.Solutions() 

	if *y != -1 {
		solutionsToPrint = filter(solutionsToPrint, func(s utils.Solution) bool {return s.Year == *y})
	}
	if *d != -1 {
		solutionsToPrint = filter(solutionsToPrint, func(s utils.Solution) bool {return s.Day == *d})
	}
	if *p != -1 {
		solutionsToPrint = filter(solutionsToPrint, func(s utils.Solution) bool {return s.Part == *p})
	}

	if *t || *q {
		passed, failed := 0, 0
		for _, r := range testSolutions(solutionsToPrint) {
			if r.err != nil {
				fmt.Printf("[FAIL] - %s: %v\n", r.solution.Name(), r.err)
				failed++
			} else {
				if (!*q) {
					fmt.Printf("[PASS] - %s\n", r.solution.Name())
				}
				passed++
			}
		}
		fmt.Printf("Passed: %d - Failed: %d\n", passed, failed)
	} else {
		for _, s := range solutionsToPrint {
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

func filter(sols []utils.Solution, filterFunc func(utils.Solution) bool) []utils.Solution {
	newSols := []utils.Solution{}
	for _, sol := range sols {
		if filterFunc(sol) {
			newSols = append(newSols, sol)
		}
	}
	return newSols
}