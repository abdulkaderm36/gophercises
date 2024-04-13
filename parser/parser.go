package parser

import (
	"io"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

type Parser struct {
	File io.Reader
}

func (p Parser) Parse() ([]Link, error) {
	tokenizer := html.NewTokenizer(p.File)
	linksMap := make(map[string]string)

	depth := 0
	key := ""
    var err error
	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			err = tokenizer.Err()
            break
		case html.StartTagToken, html.EndTagToken:
			tn, _ := tokenizer.TagName()
			if len(tn) == 1 && tn[0] == 'a' {
				if tt == html.StartTagToken {
					depth++
				} else {
					depth--
				}

				if depth == 1 {
					tk, tv, _ := tokenizer.TagAttr()
					if string(tk) == "href" {
						key = string(tv)
					}
				} else if depth == 0 {
					key = ""
				}
			}
		case html.TextToken:
			if depth == 1 {
				b := tokenizer.Text()
				if key != "" {
					linksMap[key] += string(b)
				}
			}
		case html.CommentToken:
			continue
		}

        if err != nil {
            if err.Error() == "EOF" {
                break
            }
            return nil, err
        }
	}

	links := []Link{}
	for k, v := range linksMap {
		links = append(links, Link{
			Href: k,
			Text: v,
		})
	}

	return links, nil
}
