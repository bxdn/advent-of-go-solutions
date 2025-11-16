package generation

import (
	"advent-of-go/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func Answers(year, day int) error {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
	req, e := http.NewRequest(http.MethodGet, url, nil)
	if e != nil {
		return fmt.Errorf("error creating request: %w", e)
	}
	cookieContents, e := utils.GetFileContents("private/cookie.txt")
	if e != nil {
		return fmt.Errorf("error reading cookie file: %w", e)
	}
	c := http.Cookie{Domain: "adventofcode.com", Path: "/", Value: cookieContents, Name: "session"}
	req.AddCookie(&c)
	req.Header.Add("User-Agent", "github.com/bxdn/advent-of-go by @bxdn")
	res, e := (&http.Client{Timeout: 5 * time.Second}).Do(req)
	if e != nil {
		return fmt.Errorf("error sending request: %w", e)
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("bad status code: %d", res.StatusCode)
	}
	defer res.Body.Close()

	answers, e := ArticleParagraphCodes(res.Body)
	if e != nil || len(answers) != 2 {
		return fmt.Errorf("error extracting answers: %w", e)
	}

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
