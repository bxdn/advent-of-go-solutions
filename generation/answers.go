package generation

import (
	"advent-of-go/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func Answers(year, day int) error {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
	res, e := prepareRequest(http.MethodGet, url, nil, false)
	if e != nil {
		return fmt.Errorf("error creating/sending request: %w", e)
	}
	defer res.Body.Close()
	answers, e := articleParagraphCodes(res.Body)
	if e != nil {
		return fmt.Errorf("error extracting answers: %w", e)
	}
	if len(answers) > 2 {
		return fmt.Errorf("error: more than 2 answers found for year %d day %d, this should not be possible, perhaps the html of the site changed?", year, day)
	}
	return saveAnswers(year, day, answers)
}

func saveAnswers(year, day int, answers []string) error {
	var answersInFile map[string]any
	if e := json.Unmarshal(utils.Unpack(os.ReadFile("private/answers.json")), &answersInFile); e != nil {
		return fmt.Errorf("error unmarshaling answers file: %w", e)
	}

	yearAnswers, found := answersInFile[fmt.Sprintf("%d", year)]
	if !found {
		yearAnswers = map[string]any{}
		answersInFile[fmt.Sprintf("%d", year)] = yearAnswers
	}
	yearAnswers.(map[string]any)[fmt.Sprintf("%d", day)] = answers

	answersFile, e := os.Create("private/answers.json")
	if e != nil {
		return fmt.Errorf("error creating answers file: %w", e)
	}
	defer answersFile.Close()
	_, e = answersFile.Write(utils.Unpack(json.MarshalIndent(answersInFile, "", "  ")))
	return e
}
