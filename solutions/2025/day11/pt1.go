package day11

import (
	"advent-of-go/utils"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        11,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	graph := map[string][]string{}
	for _, line := range utils.GetLines(input) {
		addLine(line, graph)
	}
	return strconv.Itoa(dfs("you", "out", graph, map[string]int{})), nil
}

func dfs(key, target string, graph map[string][]string, memo map[string]int) int {
	if val, ok := memo[key]; ok {
		return val
	}
	if key == target {
		return 1
	}
	count := 0
	for _, val := range graph[key] {
		count += dfs(val, target, graph, memo)
	}
	memo[key] = count
	return count
}

func addLine(line string, graph map[string][]string) {
	from := line[:3]
	to := strings.Split(strings.TrimSpace(line[4:]), " ")
	graph[from] = to
}
