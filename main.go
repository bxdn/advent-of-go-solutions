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
	g := flag.Bool("g", false, "Use to generate new solution set, needs year and day flags to work")
	i := flag.Bool("i", false, "Use to retrieve input and place it in the correct spot, needs year and day flags to work")
	a := flag.Bool("a", false, "Use to retrieve answers and place them in the correct spot, needs year and day flags to work")
	s := flag.Bool("s", false, "Use to submit a solution, needs year, day and part flags to work")

	flag.Parse()

	if e := runInit(); e != nil {
		fmt.Printf("Error initializing: %v\n", e)
		return
	}

	if *g || *i || *a {
		handleGeneration(g, i, a, y, d)
		return
	}
	filteredSolutions := getFilteredSolutions(y, d, p)
	if *s {
		handleSubmission(y, d, p, filteredSolutions)
	} else if *t || *q {
		handleTesting(filteredSolutions, q)
	} else {
		for _, s := range filteredSolutions {
			printSolution(s)
		}
	}
}

func handleSubmission(y, d, p *int, solutions []utils.Solution) {
	if *y == -1 || *d == -1 || *p == -1 {
		flag.PrintDefaults()
		return
	}
	if len(solutions) != 1 {
		fmt.Printf("Error: expected exactly one solution to submit, but found %d\n", len(solutions))
		return
	}
	solution := solutions[0]
	msg, e := generation.Submit(*y, *d, *p, solution)
	if e != nil {
		fmt.Printf("Error submitting solution: %v\n", e)
		return
	}
	fmt.Printf("%s Response: %s\n", solution.Name(), msg)
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
		if e := generation.Generate(*y, *d); e != nil {
			fmt.Printf("Error generating solution: %v\n", e)
		}
	}
	if *i {
		if e := generation.Input(*y, *d); e != nil {
			fmt.Printf("Error retrieving input: %v\n", e)
		}
	}
	if *a {
		if e := generation.Answers(*y, *d); e != nil {
			fmt.Printf("Error retrieving answers: %v\n", e)
		}
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
