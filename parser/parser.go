package parser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

type Parser struct {
	File io.Reader
}

func linkNodes(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}

	var ret []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}

	return ret
}

func textNodes(node *html.Node) []string {
	if node.Type == html.TextNode {
		return []string{node.Data}
	}

	var ret []string
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, textNodes(c)...)
	}

	return ret
}

func buildLink(n *html.Node) Link {
	var link Link
	for _, a := range n.Attr {
		if a.Key == "href" {
			link.Href = a.Val
			break
		}
	}
	link.Text = strings.TrimSpace(strings.Join(textNodes(n), ""))
	return link
}

func (p Parser) Parse() ([]Link, error) {
	node, err := html.Parse(p.File)
	if err != nil {
		return nil, err
	}

	links := []Link{}
	for _, v := range linkNodes(node) {
		links = append(links, buildLink(v))
	}
	return links, nil
}
