package day3

import (
	"advent-of-go/utils"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2024,
		Day:        3,
		Part:       2,
		Calculator: pt2,
	}
}

var activeRe = regexp.MustCompile(`don't\(\)|mul\((\d+),(\d+)\)`)
var inactiveRe = regexp.MustCompile(`do\(\)`)

type state struct {
	total  int
	active bool
	input  string
}

func pt2(input string) (string, error) {
	s := &state{0, true, input}
	for {
		match := getMatch(s.active, s.input)
		if match == nil {
			return strconv.Itoa(s.total), nil
		}
		if e := handleMatch(s, match); e != nil {
			fmt.Errorf("Error Processingm match: %w", e)
		}
	}
}

func getMatch(active bool, input string) []int {
	var re *regexp.Regexp
	if active {
		re = activeRe
	} else {
		re = inactiveRe
	}
	return re.FindStringSubmatchIndex(input)
}

func handleMatch(s *state, match []int) error {
	switch s.input[match[0]:match[1]] {
	case "don't()":
		s.active = false
	case "do()":
		s.active = true
	default:
		if len(match) != 6 {
			return errors.New("Match somehow does not have the right number of groups")
		}
		prod, e := getProduct(s.input[match[2]:match[3]], s.input[match[4]:match[5]])
		if e != nil {
			return fmt.Errorf("Error Parsing mul() command %s: %w", s.input[match[0]:match[1]], e)
		}
		s.total += prod
	}
	s.input = s.input[match[1]:]
	return nil
}
