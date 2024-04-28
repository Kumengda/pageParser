package parser

import "golang.org/x/net/html"

type BaseExtract struct {
}

func NewBaseExtract() BaseExtract {
	return BaseExtract{}
}

func (b *BaseExtract) findClickableElements(n *html.Node, clickableElements *[]*html.Node, tagName []string) {
	if n.Type == html.ElementNode {
		for _, name := range tagName {
			if name == n.Data {
				*clickableElements = append(*clickableElements, n)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		b.findClickableElements(c, clickableElements, tagName)
	}
}
