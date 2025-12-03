package generation

import (
	"advent-of-go/utils"
	"bytes"
	"fmt"
	"net/http"
)

func Submit(year, day, part int, solution utils.Solution) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day)
	answer, _, e := solution.Calculate()
	if e != nil {
		return "", fmt.Errorf("error calculating solution: %w", e)
	}
	body := fmt.Sprintf("level=%d&answer=%s", part, answer)
	res, e := prepareRequest(http.MethodPost, url, bytes.NewBuffer([]byte(body)), true)
	if e != nil {
		return "", fmt.Errorf("error creating/sending request: %w", e)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return "", fmt.Errorf("error: submission not found for year %d day %d, perhaps the day has not been released yet?", year, day)
		}
		return "", fmt.Errorf("error: bad status code: %d", res.StatusCode)
	}
	msg, e := articleParagraphText(res.Body)
	if e != nil {
		return "", fmt.Errorf("error processing response body: %w", e)
	}
	return msg, nil
}
