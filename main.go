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

	flag.Parse()

	for _, solutionSet := range solutions.Solutions() {
		solutionsToPrint := solutionSet.Solutions
		if *y != -1 {
			solutionsToPrint = filter(solutionsToPrint, func(s utils.Solution) bool {return s.Year == *y})
		}
		if *d != -1 {
			solutionsToPrint = filter(solutionsToPrint, func(s utils.Solution) bool {return s.Day == *d})
		}
		if *p != -1 {
			solutionsToPrint = filter(solutionsToPrint, func(s utils.Solution) bool {return s.Part == *p})
		}
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