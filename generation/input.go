package generation

import (
	"advent-of-go/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func Input(year, day int) error {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
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
	dirName := fmt.Sprintf("private/inputs/%d", year)
	if e := os.MkdirAll(dirName, 0777); e != nil {
		return fmt.Errorf("error creating directory: %w", e)
	}
	inputFile, e := os.Create(fmt.Sprintf("%s/day%d.txt", dirName, day))
	if e != nil {
		return fmt.Errorf("error creating pt1 file: %w", e)
	}
	defer inputFile.Close()
	_, e = io.Copy(inputFile, res.Body)
	return e
}
