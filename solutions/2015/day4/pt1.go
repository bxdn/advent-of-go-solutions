package day4

import (
	"advent-of-go/utils"
	"crypto/md5"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2015,
		Day:        4,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	var it int64 = 1
	keyBytes := []byte(input)
	for {
		message := strconv.AppendInt(keyBytes, it, 10)
		sum := md5.Sum(message)
		if sum[0] == 0 && sum[1] == 0 && sum[2]>>4 == 0 {
			return strconv.Itoa(int(it)), nil
		}
		it++
	}
}
