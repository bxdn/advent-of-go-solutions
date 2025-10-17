package day12

import (
	"advent-of-go/utils"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type State int

const (
	Off State = iota
	On
	Unknown
)

type memoKey struct{
	states string
	layout string
}

var memo map[memoKey]int = map[memoKey]int{}

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2023,
		Day: 12,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := utils.GetLines(input)
	total := 0
	for _, line := range lines {
		states, layout, e := processLine(line)
		if e != nil {
			return "", fmt.Errorf("Error processing line: %w", e)
		}
		states = append(states, Off)
		numWays := calculateNumWays(states, layout)
		total += numWays
	}
	return strconv.Itoa(total), nil
}

func calculateNumWays(states []State, layout []int) int {
	serialized := serialize(states, layout)
	if solution, ok := memo[serialized]; ok {
		return solution
	}

	var toRet int

	if len(layout) == 0 && slices.Index(states, On) == -1 {
		toRet = 1
	} else if len(states) == 0 || len(layout) == 0 {
		toRet = 0
	} else {
		switch states[0] {
			case Off: toRet = calcOff(states, layout)
			case On: toRet = calcOn(states, layout)
			case Unknown: toRet = calcOff(states, layout) + calcOn(states, layout)
		}
	}
	memo[serialized] = toRet
	return toRet
}

func calcOff(states []State, layout []int) int {
	return calculateNumWays(states[1:], layout)
}

func calcOn(states []State, layout []int) int {
	nextGroupSize := layout[0]
	statesLen := len(states)
	// If there aren't enough states to make up the group
	if nextGroupSize > statesLen {
		return 0
	}
	// If the next after the group would be On, then the groupsize would have to be at least 1 larger
	if statesLen > nextGroupSize && states[nextGroupSize] == On {
		return 0
	}
	// If there's ever an off in the contiguous group, then it's invalid
	for i := 0; i < nextGroupSize; i++ {
		if states[i] == Off {
			return 0
		}
	}
	// Treat next state as off, continue after skipping it
	return calculateNumWays(states[nextGroupSize + 1:], layout[1:])
}

func serialize(states []State, layout []int) memoKey {
	stateSb := strings.Builder{}
	for _, i := range states {
		stateSb.WriteString(strconv.Itoa(int(i)))
	}
	layoutSb := strings.Builder{}
	for _, i := range layout {
		layoutSb.WriteString(strconv.Itoa(i))
		layoutSb.WriteRune('|')
	}
	return memoKey{stateSb.String(), layoutSb.String()}
}

func processLine(line string) ([]State, []int, error) {
	toks := strings.Split(line, " ")
	if len(toks) != 2 {
		return nil, nil, errors.New("Malformed line: Unexpected format!")
	}
	springStr, layoutStr := toks[0], toks[1]
	states := make([]State, len(springStr))
	for i, char := range springStr {
		switch char {
			case '.': states[i] = Off
			case '#': states[i] = On
			case '?': states[i] = Unknown
			default: return nil, nil, errors.New("Malformed line: unexpected character!") 
		}
	}
	layoutToks := strings.Split(layoutStr, ",")
	layout := make([]int, len(layoutToks))
	for i, digitStr := range layoutToks {
		digit, e := strconv.Atoi(digitStr)
		if e != nil {
			return nil, nil, fmt.Errorf("Malformed line: layout section contained non-digit %s: %w", digitStr, e)
		}
		layout[i] = digit
	}
	return states, layout, nil
}