package parser

import (
	"github.com/Kumengda/pageParser/resType"
	"github.com/duke-git/lancet/v2/slice"
	"golang.org/x/net/html"
	"strings"
)

type TagExtract struct {
	BaseExtract
	Tags []Tag
}

func NewTagExtract() *TagExtract {
	return &TagExtract{
		BaseExtract: NewBaseExtract(),
	}
}
func (h *TagExtract) InitTags(tags []Tag) {
	h.Tags = tags
}
func (h *TagExtract) AddTag(tagName string, attr ...string) {
	for i := 0; i < len(h.Tags); i++ {
		if h.Tags[i].TagName == tagName {
			h.Tags[i].TagAttrName = slice.Unique(append(h.Tags[i].TagAttrName, attr...))
			return
		}
	}
	h.Tags = append(h.Tags, Tag{
		TagName:     tagName,
		TagAttrName: attr,
	})
}
func (h *TagExtract) DeleteTag(tagName ...string) {
	for _, t := range tagName {
		for i := 0; i < len(h.Tags); i++ {
			if h.Tags[i].TagName == t {
				h.Tags = slice.DeleteAt(h.Tags, i)
				break
			}
		}
	}
}
func (h *TagExtract) ReSetTag(tagName string, attr ...string) {
	for i := 0; i < len(h.Tags); i++ {
		if h.Tags[i].TagName == tagName {
			h.Tags[i].TagAttrName = attr
			return
		}
	}
	h.Tags = append(h.Tags, Tag{
		TagName:     tagName,
		TagAttrName: attr,
	})
}
func (h *TagExtract) AddTagFilter(tagName string, filter func(attr []html.Attribute) bool) {
	for i := 0; i < len(h.Tags); i++ {
		if h.Tags[i].TagName == tagName {
			h.Tags[i].Filter = filter
			return
		}
	}
}
func (h *TagExtract) Extract(htmlData string) []resType.TagInfo {
	var tagExtractRes []resType.TagInfo
	reader := strings.NewReader(htmlData)
	doc, err := html.Parse(reader)
	if err != nil {
		return nil
	}
	var clickableElements []*html.Node
	var extractTags []string
	for _, v := range h.Tags {
		extractTags = append(extractTags, v.TagName)
	}
	h.findClickableElements(doc, &clickableElements, extractTags)
	for _, element := range clickableElements {
		for _, t := range h.Tags {
			if t.TagName == element.Data {
				if t.Filter != nil {
					if !t.Filter(element.Attr) {
						continue
					}
				}
				for _, at := range t.TagAttrName {
					val := extractTag(element, at)
					if val != "" {
						tagExtractRes = append(tagExtractRes, resType.TagInfo{
							Name: t.TagName,
							Attr: map[string]string{at: val},
						})
					}
				}
			}
		}
	}
	return tagExtractRes
}

func getText(n *html.Node) string {
	var text string
	var getTextInternal func(*html.Node)
	getTextInternal = func(n *html.Node) {
		if n.Type == html.TextNode {
			text += n.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			getTextInternal(c)
		}
	}
	getTextInternal(n)
	return text
}
