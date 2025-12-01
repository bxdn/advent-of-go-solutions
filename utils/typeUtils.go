package utils

import (
	"fmt"
	"strconv"
)

func StringsToRunes(strings []string) [][]rune {
	toRet := [][]rune{}
	for _, s := range strings {
		toRet = append(toRet, []rune(s))
	}
	return toRet
}

func Flatten[T any](slices [][]T) []T {
	var toRet []T
	for _, s := range slices {
		toRet = append(toRet, s...)
	}
	return toRet
}

func StringsToInts(strings []string) ([]int, error) {
	nums := make([]int, len(strings))
	for i, numStr := range strings {
		seed, e := strconv.Atoi(numStr)
		if e != nil {
			return nil, fmt.Errorf("Error parsing string %s to int: %w", numStr, e)
		}
		nums[i] = seed
	}
	return nums, nil
}
