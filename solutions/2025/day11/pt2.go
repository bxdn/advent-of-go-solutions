package day11

import (
	"advent-of-go/utils"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        11,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	graph := map[string][]string{}
	for _, line := range utils.GetLines(input) {
		addLine(line, graph)
	}
	svr2dac := dfs("svr", "dac", graph, map[string]int{})
	svr2fft := dfs("svr", "fft", graph, map[string]int{})
	dac2fft := dfs("dac", "fft", graph, map[string]int{})
	fft2dac := dfs("fft", "dac", graph, map[string]int{})
	fft2out := dfs("fft", "out", graph, map[string]int{})
	dac2out := dfs("dac", "out", graph, map[string]int{})
	total := svr2dac*dac2fft*fft2out + svr2fft*fft2dac*dac2out
	return strconv.Itoa(total), nil
}
