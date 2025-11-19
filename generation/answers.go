package generation

import (
	"advent-of-go/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func AllAnswers(solutions []utils.Solution) (bool, error) {

	var answersInFile map[string]map[string][]string
	if e := json.Unmarshal(utils.Unpack(os.ReadFile("private/answers.json")), &answersInFile); e != nil {
		return false, fmt.Errorf("error unmarshaling answers file: %w", e)
	}

	ok := true
	for _, s := range solutions {
		if e := answers(s.Year, s.Day, answersInFile); e != nil {
			fmt.Printf("Error retrieving answers for year %d day %d: %v\n", s.Year, s.Day, e)
			ok = false
		}
	}

	answersFile, e := os.Create("private/answers.json")
	if e != nil {
		return false, fmt.Errorf("error creating answers file: %w", e)
	}
	defer answersFile.Close()
	_, e = answersFile.Write(utils.Unpack(json.MarshalIndent(answersInFile, "", "  ")))
	if e != nil {
		return false, fmt.Errorf("error writing answers file: %w", e)
	}
	return ok, nil
}

func answers(year, day int, answersInFile map[string]map[string][]string) error {
	// cache step
	yearAnswers, found := answersInFile[fmt.Sprintf("%d", year)]
	if !found {
		yearAnswers = map[string][]string{}
		answersInFile[fmt.Sprintf("%d", year)] = yearAnswers
	}
	answers, found := yearAnswers[fmt.Sprintf("%d", day)]
	if found && len(answers) == 2 {
		return nil
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
	res, e := prepareRequest(http.MethodGet, url, nil, false)
	if e != nil {
		return fmt.Errorf("error creating/sending request: %w", e)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return fmt.Errorf("error: page not found for year %d day %d, perhaps the day has not been released yet?", year, day)
		}
		return fmt.Errorf("error: bad status code: %d", res.StatusCode)
	}
	answers, e = articleParagraphCodes(res.Body)
	if e != nil {
		return fmt.Errorf("error extracting answers: %w", e)
	}
	if len(answers) > 2 {
		return fmt.Errorf("error: more than 2 answers found for year %d day %d, this should not be possible, perhaps the layout of the site changed?", year, day)
	} else if len(answers) == 0 {
		return fmt.Errorf("error: no answers found for year %d day %d, have you submitted a solution?", year, day)
	}
	yearAnswers[fmt.Sprintf("%d", day)] = answers
	return nil
}
