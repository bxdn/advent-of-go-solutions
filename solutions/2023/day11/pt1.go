package day11

import (
	"advent-of-go/utils"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2023,
		Day: 11,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	return getAnswer(input, 2), nil
}

func getAnswer(input string, mult int) string {
	runes := utils.StringsToRunes(utils.GetLines(input))
	grid := utils.GridFrom2dSlice(runes)
	rowsWithGalaxies, colsWithGalaxies := getCoordsWithGalaxies(runes)
	galaxyPositions := grid.FindAll('#')
	total := countAllShortest(galaxyPositions, rowsWithGalaxies, colsWithGalaxies, mult)
	return strconv.Itoa(total)
}

func countAllShortest(galaxyPositions []utils.Point, rowsWithGalaxies, colsWithGalaxies []bool, mult int) int {
	total := 0
	numGalaxies := len(galaxyPositions)
	for i, p1 := range galaxyPositions {
		for j := i + 1; j < numGalaxies; j++ {
			p2 := galaxyPositions[j]
			dx :=countDistanceAlongAxis(p1.X, p2.X, colsWithGalaxies, mult)
			dy := countDistanceAlongAxis(p1.Y, p2.Y, rowsWithGalaxies, mult)
			total += dx + dy
		}
	}
	return total
}

func countDistanceAlongAxis(n1, n2 int, withGalaxies []bool, mult int) int {
	total := 0
	minN, maxN := min(n1, n2), max(n1, n2)
	for i := minN; i < maxN; i++ {
		if withGalaxies[i] {
			total++
		} else {
			total += mult
		}
	}
	return total
}

func getCoordsWithGalaxies(lines [][]rune) ([]bool, []bool) {
	if len(lines) == 0 {
		return []bool{}, []bool{}
	}
	rowStatuses := make([]bool, len(lines))
	colStatuses := make([]bool, len(lines[0]))

	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				rowStatuses[i] = true
				colStatuses[j] = true
			}
		}
	}
	return rowStatuses, colStatuses
}