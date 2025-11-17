package generation

import (
	"advent-of-go/utils"
	"fmt"
	"io"
	"net/http"
	"time"
)

func prepareRequest(method string, url string, body io.Reader, isForm bool) (*http.Response, error) {
	req, e := http.NewRequest(method, url, body)
	if e != nil {
		return nil, fmt.Errorf("error creating request: %w", e)
	}
	cookieContents, e := utils.GetFileContents("private/cookie.txt")
	if e != nil {
		return nil, fmt.Errorf("error reading cookie file: %w", e)
	}
	c := http.Cookie{Value: cookieContents, Name: "session"}
	req.AddCookie(&c)
	req.Header.Add("User-Agent", "github.com/bxdn/advent-of-go by @bxdn")
	if isForm {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	res, e := (&http.Client{Timeout: 10 * time.Second}).Do(req)
	if e != nil {
		return nil, fmt.Errorf("error sending request: %w", e)
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error: bad status code: %d", res.StatusCode)
	}
	return res, nil
}
