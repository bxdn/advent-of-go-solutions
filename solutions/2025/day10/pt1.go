package day10

import (
	"advent-of-go/utils"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        10,
		Part:       1,
		Calculator: pt1,
	}
}

var regexLight = regexp.MustCompile(`\[(.+)\]`)
var regexBtn = regexp.MustCompile(`\(((?:,|\d)+)\)`)
var regexJoltage = regexp.MustCompile(`\{(.+)\}`)

type searchState struct {
	steps int
	state string
}

func pt1(input string) (string, error) {
	total := 0
	for _, line := range utils.GetLines(input) {
		target, btns, _ := parseMachine(line)
		total += solveMachine(target, btns)
	}
	return strconv.Itoa(total), nil
}

func solveMachine(target string, btns [][]int) int {
	visited := map[string]int{}
	queue := []searchState{{0, strings.Repeat(".", len(target))}}
	for len(queue) != 0 {
		state := queue[0]
		queue = queue[1:]
		if _, ok := visited[state.state]; ok {
			continue
		}
		visited[state.state] = state.steps
		if state.state == target {
			return state.steps
		}
		boolState := stringToBools(state.state)
		for _, btn := range btns {
			newBoolState := slices.Clone(boolState)
			for _, light := range btn {
				newBoolState[light] = !newBoolState[light]
			}
			queue = append(queue, searchState{state.steps + 1, boolsToString(newBoolState)})
		}
	}
	panic("Queue should not be empty!")
}

func stringToBools(str string) []bool {
	toRet := make([]bool, len(str))
	for i, char := range str {
		toRet[i] = char == '#'
	}
	return toRet
}

func boolsToString(bools []bool) string {
	sb := strings.Builder{}
	for _, b := range bools {
		if b {
			sb.WriteRune('#')
		} else {
			sb.WriteRune('.')
		}
	}
	return sb.String()
}

func parseMachine(line string) (string, [][]int, string) {
	target := regexLight.FindStringSubmatch(line)[1]
	btns := [][]int{}
	for _, match := range regexBtn.FindAllStringSubmatch(line, -1) {
		btn, e := utils.StringsToInts(strings.Split(match[1], ","))
		if e != nil {
			panic(e)
		}
		btns = append(btns, btn)
	}
	joltageString := regexJoltage.FindStringSubmatch(line)[1]
	return target, btns, joltageString
}
