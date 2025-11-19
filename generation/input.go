package generation

import (
	"advent-of-go/utils"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func AllInput(solutions []utils.Solution) bool {
	ok := true
	for _, s := range solutions {
		if e := input(s.Year, s.Day); e != nil {
			fmt.Printf("Error retrieving input for year %d day %d: %v\n", s.Year, s.Day, e)
			ok = false
		}
	}
	return ok
}

func input(year, day int) error {
	// cache step
	if _, e := os.Stat(fmt.Sprintf("private/inputs/%d/day%d.txt", year, day)); e == nil {
		return nil
	} else if !os.IsNotExist(e) {
		return fmt.Errorf("error checking if input file exists: %w", e)
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	res, e := prepareRequest(http.MethodGet, url, nil, false)
	if e != nil {
		return fmt.Errorf("error creating/sending request: %w", e)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		switch res.StatusCode {
		case http.StatusNotFound:
			return fmt.Errorf("error: input not found for year %d day %d, perhaps the day has not been released yet?", year, day)
		case http.StatusBadRequest:
			return fmt.Errorf("error: bad request for year %d day %d, perhaps your cookie is invalid?", year, day)
		default:
			return fmt.Errorf("error: bad status code: %d", res.StatusCode)
		}
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	body = bytes.TrimSpace(body)

	dirName := fmt.Sprintf("private/inputs/%d", year)
	if e := os.MkdirAll(dirName, 0777); e != nil {
		return fmt.Errorf("error creating directory: %w", e)
	}
	inputFile, e := os.Create(fmt.Sprintf("%s/day%d.txt", dirName, day))
	if e != nil {
		return fmt.Errorf("error creating pt1 file: %w", e)
	}
	defer inputFile.Close()

	_, e = inputFile.Write(body)
	return e
}
