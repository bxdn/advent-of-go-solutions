package generation

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Input(year, day int) error {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	res, e := prepareRequest(http.MethodGet, url, nil, false)
	if e != nil {
		return fmt.Errorf("error creating/sending request: %w", e)
	}
	defer res.Body.Close()
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
