package day12

import (
	"advent-of-go/utils"
	"fmt"
	"regexp"
	"strconv"
)

var regex = regexp.MustCompile(`-?\d+`)

type moon struct {
	px, py, pz,
	vx, vy, vz int
}

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 12,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	moons, e := parseMoons(input)
	if e != nil {
		return "", fmt.Errorf("Error parsing moon: %w", e)
	}
	for _ = range 1000 {
		simOnce(moons)
	}
	total := 0
	for _, moon := range moons {
		total += (abs(moon.px) + abs(moon.py) + abs(moon.pz)) * (abs(moon.vx) + abs(moon.vy) + abs(moon.vz))
	}
	return strconv.Itoa(total), nil
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func simOnce(moons []moon) {
	for j := range moons {
		for k := j + 1; k < len(moons); k++ {
			moons[j], moons[k] = simPairGravity(moons[j], moons[k])
		}
	}
	for i, moon := range moons {
		moons[i].px += moon.vx
		moons[i].py += moon.vy
		moons[i].pz += moon.vz
	}
}

func simPairGravity(moon1, moon2 moon) (moon, moon) {
	if moon1.px < moon2.px {
		moon1.vx += 1
		moon2.vx -= 1
	} else if moon1.px > moon2.px {
		moon1.vx -= 1
		moon2.vx += 1
	}
	if moon1.py < moon2.py {
		moon1.vy += 1
		moon2.vy -= 1
	} else if moon1.py > moon2.py {
		moon1.vy -= 1
		moon2.vy += 1
	}
	if moon1.pz < moon2.pz {
		moon1.vz += 1
		moon2.vz -= 1
	} else if moon1.pz > moon2.pz {
		moon1.vz -= 1
		moon2.vz += 1
	}
	return moon1, moon2
}

func parseMoons(input string) ([]moon, error) {
	moonLines := utils.GetLines(input)
	moons := []moon{}
	for _, moonStr := range moonLines {
		moon, e := parseMoon(moonStr)
		if e != nil {
			return nil, fmt.Errorf("Error parsing moon: %w", e)
		}
		moons = append(moons, moon)
	}
	return moons, nil
}

func parseMoon(moonStr string) (moon, error) {
	positionStrs := regex.FindAllString(moonStr, -1)
	var m moon
	if len(positionStrs) != 3 {
		return m, fmt.Errorf("Malformed data: line %s expected to have 3 position values, had %d", moonStr, len(positionStrs))
	}
	positions, e := utils.StringsToInts(positionStrs)
	if e != nil {
		return m, fmt.Errorf("Error parsing position strings to numbers: %w", e)
	}
	return moon{px: positions[0], py: positions[1], pz: positions[2]}, nil
}