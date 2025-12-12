package day12

import (
	"advent-of-go/utils"
	"regexp"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        12,
		Part:       1,
		Calculator: pt1,
	}
}

type region struct {
	dims    utils.Point
	targets []int
}

var regex1 = regexp.MustCompile(`(?:#|\.)+`)
var regex2 = regexp.MustCompile(`(\d+)x(\d+):(.+)`)

func pt1(input string) (string, error) {
	shapeAreas := getShapeAreas(input)
	objectives := getRegions(input)
	total := 0
	for _, reg := range objectives {
		if evalRegion(reg, shapeAreas) {
			total++
		}
	}
	return strconv.Itoa(total), nil
}

func evalRegion(reg region, shapeAreas []int) bool {
	totalSpacesTaken := 0
	totalShapes := 0
	for i, target := range reg.targets {
		totalSpacesTaken += shapeAreas[i] * target
		totalShapes += target
	}

	spaces := reg.dims.X * reg.dims.Y
	if totalSpacesTaken > spaces {
		return false
	}

	shapeSlots := (reg.dims.X / 3) * (reg.dims.Y / 3)
	if shapeSlots >= totalShapes {
		return true
	}

	panic("greedy didn't work!")
}

func getShapeAreas(input string) []int {
	shapeLines := regex1.FindAllString(input, -1)
	areas := []int{}
	for i := 2; i < len(shapeLines); i += 3 {
		area := 0
		for _, line := range shapeLines[i-2 : i+1] {
			for _, char := range line {
				if char == '#' {
					area++
				}
			}
		}
		areas = append(areas, area)
	}
	return areas
}

func getRegions(input string) []region {
	regionLines := regex2.FindAllStringSubmatch(input, -1)
	regions := []region{}
	for _, line := range regionLines {
		width, e := strconv.Atoi(line[1])
		if e != nil {
			panic(e)
		}
		height, e := strconv.Atoi(line[2])
		if e != nil {
			panic(e)
		}
		targets, e := utils.StringsToInts(strings.Split(strings.TrimSpace(line[3]), " "))
		if e != nil {
			panic(e)
		}
		regions = append(regions, region{utils.Point{X: width, Y: height}, targets})
	}
	return regions
}
