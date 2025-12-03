package utils

import (
	"fmt"
	"time"
)

type Calculator func(string) (string, error)

type Solution struct {
	Year, Day, Part int
	Calculator      Calculator
}

func (s Solution) Name() string {
	return fmt.Sprintf("%d Day %d Pt %d", s.Year, s.Day, s.Part)
}

func (s Solution) Calculate() (string, string, error) {
	inputPath := fmt.Sprintf("private/inputs/%d/day%d.txt", s.Year, s.Day)
	contents, e := GetFileContents(inputPath)
	if e != nil {
		return "", "", fmt.Errorf("error getting contents of input file for %s: %w", s.Name(), e)
	}
	start := time.Now()
	answer, e := s.Calculator(contents)
	return answer, time.Since(start).String(), e
}
