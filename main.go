package main

import (
	"advent-of-go/generation"
	"advent-of-go/solutions"
	"advent-of-go/utils"
	"flag"
	"fmt"
	"time"
)

func main() {
	y := flag.Int("y", -1, "Year of solutions to display")
	d := flag.Int("d", -1, "Day of solutions to display")
	p := flag.Int("p", -1, "Part of solutions to display")
	t := flag.Bool("t", false, "Use to only test against known answers")
	q := flag.Bool("q", false, "Use to test in quiet mode (only failures logged)")
	g := flag.Bool("g", false, "Use to generate new solution set, needs year and day flags to work")
	s := flag.Bool("s", false, "Use to submit a solution, needs year, day and part flags to work")
	n := flag.Bool("n", false, "Use to use the current day and year for the year and day flags (only works in December)")

	flag.Parse()

	if e := processNowFlag(y, d, n); e != nil {
		fmt.Printf("%v\n", e)
		return
	}

	if e := runInit(); e != nil {
		fmt.Printf("Error initializing: %v\n", e)
		return
	}

	if *g {
		handleGeneration(y, d)
		return
	}

	solutions := solutions.Solutions()
	if len(solutions) == 0 {
		fmt.Println("No solutions found! Are you sure you have implemented any solutions?")
		return
	}
	filteredSolutions := getFilteredSolutions(solutions, y, d, p)
	if len(filteredSolutions) == 0 {
		println("No solutions found matching your criteria! Perhaps your filter is too strict. (Are you using incorrect -y -d -p -n flags?)")
		return
	}

	if ok := generation.AllInput(solutions); !ok {
		fmt.Printf("Finished pulling input with errors!\n")
	}

	// standard print command if no other flags are set
	if !*s && !*t && !*q {
		for _, s := range filteredSolutions {
			printSolutionResults(s)
		}
		return
	}

	if *s {
		handleSubmission(y, d, p, filteredSolutions)
	}
	if *t || *q {
		handleTesting(filteredSolutions, q)
	}
}

func processNowFlag(y, d *int, n *bool) error {
	if !*n {
		return nil
	}
	year, month, day := time.Now().Date()
	if month != time.December {
		return fmt.Errorf("error: -n flag can only be used in December")
	}
	*y = year
	*d = day
	return nil
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
	if ok, e := generation.AllAnswers(solutions); e != nil {
		fmt.Printf("Error retrieving answers: %v\n", e)
		return
	} else if !ok {
		fmt.Printf("Finished pulling answers with errors!\n")
	}
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

func handleGeneration(y, d *int) {
	if *y == -1 || *d == -1 {
		fmt.Println("Error: -g flag requires -y and -d flags to be set (or -n)")
	} else {
		if e := generation.Generate(*y, *d); e != nil {
			fmt.Printf("Error generating solution: %v\n", e)
			return
		}
		fmt.Printf("Successfully generated solution for year %d day %d\n", *y, *d)
	}
}

func getFilteredSolutions(solutions []utils.Solution, y, d, p *int) []utils.Solution {
	if *y != -1 {
		solutions = filter(solutions, func(s utils.Solution) bool { return s.Year == *y })
	}
	if *d != -1 {
		solutions = filter(solutions, func(s utils.Solution) bool { return s.Day == *d })
	}
	if *p != -1 {
		solutions = filter(solutions, func(s utils.Solution) bool { return s.Part == *p })
	}
	return solutions
}

func printSolutionResults(s utils.Solution) {
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
