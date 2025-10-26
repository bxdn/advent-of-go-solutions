package day14

import (
	"advent-of-go/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 14,
		Part: 1,
		Calculator: pt1,
	}
}

var regex = regexp.MustCompile(`\d+ [A-Z]+`)

type process struct {
	out int
	in []resourcePair
}

type resourcePair struct {
	number int
	resource string
}

func pt1(input string) (string, error) {
	costs, e := getCosts(input)
	if e != nil {
		return "", fmt.Errorf("Error getting costs: %w", e)
	}
	total := get1FuelTotal(costs)
	return strconv.Itoa(total), nil
}

func get1FuelTotal(costs map[string]process) int {
	needed := []resourcePair{{1, "FUEL"}}
	remaining := map[string]int{}
	total := 0
	for len(needed) > 0 {
		var numOre int
		numOre, needed = processItem(needed, remaining, costs)
		total += numOre
	}
	return total
}

func getCosts(input string) (map[string]process, error) {
	lines := utils.GetLines(input)
	costs := map[string]process{}
	for _, line := range lines {
		in, out, e := parseLine(line)
		if e != nil {
			return nil, fmt.Errorf("Error parsing line: %w", e)
		}
		costs[out] = in
	}
	return costs, nil
}

func parseLine(line string) (process, string, error) {
	var proc process
	pairStrs := regex.FindAllString(line, -1)
	if len(pairStrs) < 2 {
		return proc, "", fmt.Errorf("Malformed line %s: expected at least 2 items", line)
	}
	pairs := make([]resourcePair, len(pairStrs))
	for i, pairStr := range pairStrs {
		pair, e := parsePair(pairStr)
		if e != nil {
			return proc, "", fmt.Errorf("Error while parsing pair %s: %w", pairStr, e)
		}
		pairs[i] = pair
	}
	proc.in = pairs[:len(pairs) - 1]
	out := utils.Last(pairs).OrPanic("Somehow pairs are empty even though we guaranteed at least 2")
	proc.out = out.number
	return proc, out.resource, nil
}

func parsePair(item string) (resourcePair, error) {
	var pair resourcePair
	nStr, resource, found := strings.Cut(item, " ")
	if !found {
		return pair, fmt.Errorf("Malformed line %s: expected space in output but was absent", item)
	}
	n, e := strconv.Atoi(nStr)
	if e != nil {
		return pair, fmt.Errorf("Error parsing input amount: %w", e)
	}
	return resourcePair{n, resource}, nil
}

func processItem(queue []resourcePair, remaining map[string]int, costs map[string]process) (int, []resourcePair) {
	needed := adjustNeeded(queue[0], remaining)
	queue = queue[1:]
	if needed.number == 0 {
		return 0, queue
	}
	if needed.resource == "ORE" {
		return needed.number, queue
	}
	proc := costs[needed.resource]
	numRecipesMade := needed.number / proc.out
	if needed.number % proc.out != 0 || needed.number < proc.out {
		numRecipesMade++
	}
	rem := proc.out * numRecipesMade - needed.number
	if rem != 0 {
		remaining[needed.resource] += rem
	}
	for _, input := range proc.in {
		numInputNeeded := numRecipesMade * input.number
		if numInputNeeded != 0 {
			queue = append(queue, resourcePair{numInputNeeded, input.resource})
		}
	}
	return 0, queue
}

func adjustNeeded(needed resourcePair, remaining map[string]int) resourcePair {
	if v, ok := remaining[needed.resource]; ok {
		if v > needed.number {
			remaining[needed.resource] -= needed.number
			needed.number = 0
		} else {
			needed.number -= v
			delete(remaining, needed.resource)
		}
	}
	return needed
}