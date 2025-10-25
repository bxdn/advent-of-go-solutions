package day4

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

type RepeatState int
const (
	None RepeatState = iota // Current char hasn't repeated
	Once // We found one repetition of current char
	More // We found more than 2 repetitions of current char
	Locked // We found exactly 2 repetitions at some point
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 4,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	inputRange, e := utils.StringsToInts(strings.Split(input, "-"))
	if e != nil {
		return "", fmt.Errorf("Error parsing tokens: %w", e)
	}
	if len(inputRange) != 2 {
		return "", fmt.Errorf("Input %s is malformed: should have 2 - delimited sections", input)
	}
	min, max := inputRange[0], inputRange[1]
	total := 0
	for i := min; i <= max; i++ {
		if isValidPassword2(i) {
			total++
		}
	}
	return strconv.Itoa(total), nil
}

func isValidPassword2(n int) bool {
	prevRune := '.'
	nStr := strconv.Itoa(n)
	repeat := None
	for _, char := range nStr {
		if char < prevRune {
			return false
		}
		repeat = getRepeatState(char, prevRune, repeat)
		prevRune = char
	}
	return repeat == Once || repeat == Locked
}

func getRepeatState(char, prevChar rune, state RepeatState) RepeatState {
	if state == Locked {
		return Locked
	}
	if char == prevChar {
		if state == None {
			return Once
		} else {
			return More
		}
	} else {
		if state == Once {
			return Locked
		} else {
			return None
		}
	}
}