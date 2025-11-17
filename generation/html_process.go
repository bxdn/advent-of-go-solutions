package generation

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

func extractText(n *html.Node) (string, error) {
	var b strings.Builder
	var f func(*html.Node) error
	f = func(n *html.Node) error {
		if n.Type == html.TextNode {
			_, e := b.WriteString(n.Data)
			if e != nil {
				return e
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if e := f(c); e != nil {
				return e
			}
		}
		return nil
	}
	f(n)
	return strings.TrimSpace(b.String()), nil
}

func nextElementSibling(n *html.Node) *html.Node {
	for s := n.NextSibling; s != nil; s = s.NextSibling {
		if s.Type == html.ElementNode {
			return s
		}
	}
	return nil
}

func articleParagraphCodes(r io.Reader) ([]string, error) {
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
		answer := potentialCode.FirstChild.Data
		if e != nil {
			return e
		}
		results = append(results, strings.TrimSpace(answer))
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

func articleParagraphText(r io.Reader) (string, error) {
	doc, e := html.Parse(r)
	if e != nil {
		return "", e
	}

	var walk func(*html.Node) (string, error)
	walk = func(n *html.Node) (string, error) {
		if n.Type == html.ElementNode && n.Data == "article" {
			return extractText(n.FirstChild)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if s, e := walk(c); e != nil {
				return "", e
			} else if s != "" {
				return s, nil
			}
		}
		return "", nil
	}

	return walk(doc)
}
