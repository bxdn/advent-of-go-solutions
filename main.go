package main

import (
	"advent-of-go/generation"
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
	q := flag.Bool("q", false, "Use to test in quiet mode (only failures logged)")
	g := flag.Bool("g", false, "Use to generate new solution set")
	i := flag.Bool("i", false, "Use to retrieve input and place it in the correct spot")
	a := flag.Bool("a", false, "Use to retrieve answers and place them in the correct spot")

	flag.Parse()

	if *g || *i || *a {
		handleGeneration(g, i, a, y, d)
		return
	}
	solutionsToPrint := getFilteredSolutions(y, d, p)
	if *t || *q {
		handleTesting(solutionsToPrint, q)
	} else {
		for _, s := range solutionsToPrint {
			printSolution(s)
		}
	}
}

func handleTesting(solutions []utils.Solution, q *bool) {
	passed, failed := 0, 0
	for _, r := range testSolutions(solutions) {
		if r.err != nil {
			fmt.Printf("[FAIL] - %s: %v\n", r.solution.Name(), r.err)
			failed++
		} else {
			if !*q {
				fmt.Printf("[PASS] - %s\n", r.solution.Name())
			}
			passed++
		}
	}
	fmt.Printf("Passed: %d - Failed: %d\n", passed, failed)
}

func handleGeneration(g, i, a *bool, y, d *int) {
	if *y == -1 || *d == -1 {
		flag.PrintDefaults()
		return
	}
	if *g {
		utils.Must(generation.Generate(*y, *d))
	}
	if *i {
		utils.Must(generation.Input(*y, *d))
	}
	if *a {
		utils.Must(generation.Answers(*y, *d))
	}
}

func getFilteredSolutions(y, d, p *int) []utils.Solution {
	solutionsToPrint := solutions.Solutions()
	if *y != -1 {
		solutionsToPrint = filter(solutionsToPrint, func(s utils.Solution) bool { return s.Year == *y })
	}
	if *d != -1 {
		solutionsToPrint = filter(solutionsToPrint, func(s utils.Solution) bool { return s.Day == *d })
	}
	if *p != -1 {
		solutionsToPrint = filter(solutionsToPrint, func(s utils.Solution) bool { return s.Part == *p })
	}
	return solutionsToPrint
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
	var newSols []utils.Solution
	for _, sol := range sols {
		if filterFunc(sol) {
			newSols = append(newSols, sol)
		}
	}
	return newSols
}
