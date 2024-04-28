package parser

import "golang.org/x/net/html"

type Tag struct {
	TagName     string
	TagAttrName []string
	Filter      func(attr []html.Attribute) bool
}

// https://www.w3cschool.cn/htmltags
//
//head:profile,frame,iframe,img:longdesc,meta:scheme;object:archive,codebase,data,html:manifest;source:src;track:src,poster:src
var DefaultTagRules = []Tag{
	{
		TagName:     "a",
		TagAttrName: []string{"href"},
	},
	{
		TagName:     "area",
		TagAttrName: []string{"href"},
	},
	{
		TagName:     "base",
		TagAttrName: []string{"href"},
	},
	{
		TagName:     "link",
		TagAttrName: []string{"href"},
	},
	{
		TagName:     "blockquote",
		TagAttrName: []string{"cite"},
	},
	{
		TagName:     "del",
		TagAttrName: []string{"cite"},
	},
	{
		TagName:     "ins",
		TagAttrName: []string{"cite"},
	},
	{
		TagName:     "q",
		TagAttrName: []string{"cite"},
	},
	{
		TagName:     "body",
		TagAttrName: []string{"background"},
	},
	{
		TagName:     "button",
		TagAttrName: []string{"formaction"},
	},
	{
		TagName:     "embed",
		TagAttrName: []string{"src"},
	},
	{
		TagName:     "frame",
		TagAttrName: []string{"src"},
	},
	{
		TagName:     "iframe",
		TagAttrName: []string{"src"},
	},
	{
		TagName:     "img",
		TagAttrName: []string{"src"},
	},
	{
		TagName:     "input",
		TagAttrName: []string{"src"},
	},
	{
		TagName:     "script",
		TagAttrName: []string{"src"},
	},
}
