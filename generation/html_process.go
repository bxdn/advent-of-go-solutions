package generation

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/net/html"
)

func innerHTML(n *html.Node) (string, error) {
	var buf bytes.Buffer
	if e := html.Render(&buf, n.FirstChild); e != nil {
		return "", e
	}
	return buf.String(), nil
}

func nextElementSibling(n *html.Node) *html.Node {
	for s := n.NextSibling; s != nil; s = s.NextSibling {
		if s.Type == html.ElementNode {
			return s
		}
	}
	return nil
}

func ArticleParagraphCodes(r io.Reader) ([]string, error) {
	doc, e := html.Parse(r)
	if e != nil {
		return nil, e
	}

	var results []string

	findCode := func(x *html.Node) error {
		potentialCode := x.FirstChild.NextSibling
		if potentialCode == nil || potentialCode.Data != "code" {
			return nil
		}
		innerHtml, e := innerHTML(potentialCode)
		if e != nil {
			return e
		}
		results = append(results, strings.TrimSpace(innerHtml))
		return nil
	}

	var walk func(*html.Node) error
	walk = func(n *html.Node) error {
		if n.Type == html.ElementNode && n.Data == "article" {
			p := nextElementSibling(n)
			if e := findCode(p); e != nil {
				return e
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if e := walk(c); e != nil {
				return e
			}
		}
		return nil
	}

	e = walk(doc)
	return results, e
}
