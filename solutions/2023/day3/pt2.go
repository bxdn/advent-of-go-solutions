package day3

import (
	"advent-of-go/utils"
	"fmt"
	"regexp"
	"strconv"
)

var starRegex = regexp.MustCompile(`\*`)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2023, 
		Day: 3,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := utils.GetLines(input)
	numGrid, e := initializeNumGrid(lines)
	if e != nil {
		return "", fmt.Errorf("Error initializing the number grid: %w", e)
	}

	total := 0
	for y, line := range lines {
		num, e := getLineTotal(y, line, numGrid)
		if e != nil {
			return "", fmt.Errorf("Error getting the total for the line: %w", e)
		}
		total += num
	}
	
	return strconv.Itoa(total), nil
}

func initializeNumGrid(lines []string) (utils.FinGrid[numWithPos], error) {
	width, height := 0, len(lines)
	if height > 0 {
		width = len(lines[0])
	}

	nums := []numWithPos{}
	for y, line := range lines {
		nums = append(nums, getLineNums(y, line)...)
	}

	grid := utils.BlankGrid[numWithPos](width, height)
	for _, num := range nums {
		for x := num.x; x < num.x + len(num.numString); x++ {
			if e := grid.Set(x, num.y, num); e != nil {
				return grid, fmt.Errorf("Error setting the num at the position: %w", e)
			}
		}
	}

	return grid, nil
}

func getLineTotal(y int, line string, grid utils.FinGrid[numWithPos]) (int, error) {
	total := 0
	for _, match := range starRegex.FindAllStringIndex(line, -1) {
		ratio, e := getTotalForNumMap(getNumMapForGear(match[0], y, grid))
		if e != nil {
			return 0, fmt.Errorf("Error getting the total for gear ratio: %w", e)
		} 
		total += ratio
	}
	return total, nil
}

func getNumMapForGear(x, y int, grid utils.FinGrid[numWithPos]) map[numWithPos]string {
	nums := map[numWithPos]string{}
	for xOffset := -1; xOffset <= 1; xOffset++ {
		for yOffset := -1; yOffset <= 1; yOffset++ {
			num := grid.At(x + xOffset, y + yOffset).Or(numWithPos{})
			if len(num.numString) != 0 {
				nums[num] = num.numString
			}
		}
	}
	return nums
}

func getTotalForNumMap(nums map[numWithPos]string) (int, error) {
	if len(nums) != 2 {
		return 0, nil
	}
	total := 1
	for _, str := range nums {
		intValue, e := strconv.Atoi(str)
		if e != nil {
			fmt.Errorf("Error parsing the numWithPos: %w", e)
		}
		total *= intValue
	}
	return total, nil
}