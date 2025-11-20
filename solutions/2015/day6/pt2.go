package day6

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2015,
		Day:        6,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	light_arr := make([]int, 1_000_000)
	grid := utils.GridFromSlice(light_arr, 1_000)
	for _, line := range utils.GetLines(input) {
		inst, e := parseLine(line)
		if e != nil {
			return "", fmt.Errorf("error: error parsing line: %w", e)
		}
		doInstructionImproved(grid, inst)
	}
	total := 0
	for _, light := range light_arr {
		total += light
	}
	return strconv.Itoa(total), nil
}

func doInstructionImproved(lights utils.FinGrid[int], inst instruction) error {
	for x := inst.start.X; x <= inst.end.X; x++ {
		for y := inst.start.Y; y <= inst.end.Y; y++ {
			curVal, e := lights.At(x, y).OrErr("Light should be instantiated here, bad coords?")
			if e != nil {
				return fmt.Errorf("error getting current light value: %w", e)
			}
			switch inst.typ {
			case on:
				e = lights.Set(x, y, curVal+1)
			case off:
				e = lights.Set(x, y, max(0, curVal-1))
			case toggle:
				e = lights.Set(x, y, curVal+2)
			}
			if e != nil {
				return fmt.Errorf("error doing instruction: %w", e)
			}
		}
	}
	return nil
}
