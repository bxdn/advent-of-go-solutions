package utils

import (
	"fmt"
	"regexp"
)

var intRegex = regexp.MustCompile(`-?\d+`)
var posIntRegex = regexp.MustCompile(`\d+`)
var negIntRegex = regexp.MustCompile(`-\d+`)

func FindInts(s string) []int {
	numStrs := intRegex.FindAllString(s, -1)
	nums, e := StringsToInts(numStrs)
	if e != nil {
		panic(fmt.Sprintf("Error parsing numbers %v conforming to int regex: %v", numStrs, e))
	}
	return nums
}

func FindPosInts(s string) []int {
	numStrs := posIntRegex.FindAllString(s, -1)
	nums, e := StringsToInts(numStrs)
	if e != nil {
		panic(fmt.Sprintf("Error parsing numbers %v conforming to positive int regex: %v", numStrs, e))
	}
	return nums
}

func FindNegInts(s string) []int {
	numStrs := negIntRegex.FindAllString(s, -1)
	nums, e := StringsToInts(numStrs)
	if e != nil {
		panic(fmt.Sprintf("Error parsing numbers %v conforming to negative int regex: %v", numStrs, e))
	}
	return nums
}