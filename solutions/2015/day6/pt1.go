package day6

import (
	"advent-of-go/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2015,
		Day:        6,
		Part:       1,
		Calculator: pt1,
	}
}

type instructionType int

type instruction struct {
	typ        instructionType
	start, end utils.Point
}

const (
	off instructionType = iota
	on
	toggle
)

var re = regexp.MustCompile(`\d+`)

func pt1(input string) (string, error) {
	light_arr := make([]bool, 1_000_000)
	grid := utils.GridFromSlice(light_arr, 1_000)
	for _, line := range utils.GetLines(input) {
		inst, e := parseLine(line)
		if e != nil {
			return "", fmt.Errorf("error: error parsing line: %w", e)
		}
		doInstruction(grid, inst)
	}
	total := 0
	for _, light := range light_arr {
		if light {
			total++
		}
	}
	return strconv.Itoa(total), nil
}

func doInstruction(lights utils.FinGrid[bool], inst instruction) error {
	for x := inst.start.X; x <= inst.end.X; x++ {
		for y := inst.start.Y; y <= inst.end.Y; y++ {
			var e error
			switch inst.typ {
			case off:
				e = lights.Set(x, y, false)
			case on:
				e = lights.Set(x, y, true)
			case toggle:
				{
					var curVal bool
					if curVal, e = lights.At(x, y).OrErr("Light should be instantiated here, bad coords?"); e == nil {
						e = lights.Set(x, y, !curVal)
					}
				}
			}
			if e != nil {
				return fmt.Errorf("error doing instruction: %w", e)
			}
		}
	}
	return nil
}

func parseLine(line string) (instruction, error) {
	toRet := instruction{}
	str_nums := re.FindAllString(line, -1)
	if len(str_nums) != 4 {
		return toRet, fmt.Errorf("error: expected 4 numbers in the range, got %d", len(str_nums))
	}
	// Guaranteed since regex matches any int
	start_x := utils.Unpack(strconv.Atoi(str_nums[0]))
	start_y := utils.Unpack(strconv.Atoi(str_nums[1]))
	toRet.start = utils.Point{X: start_y, Y: start_x}
	end_x := utils.Unpack(strconv.Atoi(str_nums[2]))
	end_y := utils.Unpack(strconv.Atoi(str_nums[3]))
	toRet.end = utils.Point{X: end_y, Y: end_x}

	if strings.Contains(line, "off") {
		toRet.typ = off
	} else if strings.Contains(line, "on") {
		toRet.typ = on
	} else {
		toRet.typ = toggle
	}
	return toRet, nil
}
