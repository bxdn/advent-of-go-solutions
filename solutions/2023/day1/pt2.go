package day1

import (
	"advent-of-go/utils"
	"math"
	"strconv"
	"strings"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2023, 
		Day: 1,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := utils.GetLines(input)
	lookupTable := map[string]string{
		"one": "1",
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
		"1": "1",
		"2": "2",
		"3": "3",
		"4": "4",
		"5": "5",
		"6": "6",
		"7": "7",
		"8": "8",
		"9": "9",
	}
	total := 0
	for _, line := range lines {
		num, e := getLineNumWithLetters(line, lookupTable)
		if e != nil {
			return "", e
		}
		total += num
	}
	return strconv.Itoa(total), nil
}

func getLineNumWithLetters(line string, lookup map[string]string) (int, error) {
	minKey, maxKey, minIdx, maxIdx := "", "", math.MaxInt, -1
	for k := range lookup {
		i := strings.Index(line, k)
		if i != -1 && i < minIdx {
			minKey, minIdx = k, i
		}
		i = strings.LastIndex(line, k)
		if i > maxIdx {
			maxKey, maxIdx = k, i
		}
	}
	return strconv.Atoi(lookup[minKey] + lookup[maxKey])
}