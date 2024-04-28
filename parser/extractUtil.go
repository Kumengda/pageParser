package parser

import "golang.org/x/net/html"

func extractTag(element *html.Node, tagName string) string {
	for _, attr := range element.Attr {
		if attr.Key == tagName {
			return attr.Val
		}
	}
	return ""
}
