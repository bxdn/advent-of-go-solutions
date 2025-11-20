package day4

import (
	"advent-of-go/utils"
	"crypto/md5"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2015,
		Day:        4,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	var it int64 = 1
	keyBytes := []byte(input)
	for {
		message := strconv.AppendInt(keyBytes, it, 10)
		sum := md5.Sum(message)
		if sum[0] == 0 && sum[1] == 0 && sum[2] == 0 {
			return strconv.Itoa(int(it)), nil
		}
		it++
	}
}
