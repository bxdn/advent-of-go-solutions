package day16

import (
	"advent-of-go/utils"
	"errors"
	"fmt"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 16,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	current := stringToInts(input)
	offset, err := strconv.Atoi(input[:7])

	if err != nil {
		return "", fmt.Errorf("Error converting offset number: %w", err)
	}

	totalLen := len(current) * 10000
	if offset < totalLen/2 {
		return "", errors.New("offset not in second half; suffix-sum trick won't work")
	}

	tailLen := totalLen - offset
	data := make([]int, tailLen)
	for i := 0; i < tailLen; i++ {
		data[i] = current[(offset+i)%len(current)]
	}

	for range 100 {
		sum := 0
		for i := tailLen - 1; i >= 0; i-- {
			sum = (sum + data[i]) % 10
			data[i] = sum
		}
	}

	return strconv.Itoa(concat(data[:8])), nil
}