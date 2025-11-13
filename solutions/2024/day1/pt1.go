package day1

import (
	"advent-of-go/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2024,
		Day:        1,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	num_strs := strings.Fields(input)
	leftList, rightList := []int{}, []int{}
	for i, str := range num_strs {
		num, e := strconv.Atoi(str)
		if e != nil {
			return "", fmt.Errorf("Error Parsing number: %w", e)
		}
		if i&1 == 0 {
			leftList = append(leftList, num)
		} else {
			rightList = append(rightList, num)
		}
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	total := 0
	for i := range leftList {
		total += utils.Abs(leftList[i] - rightList[i])
	}
	return strconv.Itoa(total), nil
}
