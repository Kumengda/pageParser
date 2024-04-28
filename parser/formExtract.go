package parser

import (
	"github.com/Kumengda/pageParser/resType"
	"golang.org/x/net/html"
	"strings"
)

type FormExtract struct {
	BaseExtract
}

func NewFormExtract() *FormExtract {
	return &FormExtract{
		NewBaseExtract(),
	}
}

func (f *FormExtract) Extract(htmlData string) []resType.FormDatas {
	reader := strings.NewReader(htmlData)
	doc, err := html.Parse(reader)
	if err != nil {
		return nil
	}
	var clickableElements []*html.Node
	var formDatas []resType.FormDatas
	f.findClickableElements(doc, &clickableElements, []string{"input"})
	for _, element := range clickableElements {
		formElement := isChildOfForm(element)
		if formElement == nil {
			continue
		}
		var formInfo resType.FormDatas
		for _, v := range formElement.Attr {
			switch v.Key {
			case "enctype":
				formInfo.Enctype = v.Val
			case "action":
				formInfo.Action = v.Val
			case "method":
				formInfo.Method = v.Val
			}
		}
		if !isRepeatForm(formDatas, formInfo) {
			formDatas = append(formDatas, formInfo)
		}
		var oneFormData resType.FormData
		for _, v := range element.Attr {
			switch v.Key {
			case "name":
				oneFormData.Name = v.Val
			case "type":
				oneFormData.Type = v.Val
			}
		}
		oneFormData.Value = element.Data
		for i := 0; i < len(formDatas); i++ {
			if formDatas[i].Action == formInfo.Action && formDatas[i].Method == formInfo.Method && formDatas[i].Enctype == formInfo.Enctype {
				formDatas[i].FormData = append(formDatas[i].FormData, oneFormData)
			}
		}
	}
	return formDatas
}

func isRepeatForm(formDatas []resType.FormDatas, datas resType.FormDatas) bool {
	for _, v := range formDatas {
		if v.Action == datas.Action && v.Method == datas.Method && v.Enctype == datas.Enctype {
			return true
		}
	}
	return false
}

func isChildOfForm(n *html.Node) *html.Node {
	if n.Parent == nil {
		return nil
	}
	if n.Parent.Data == "form" {
		return n.Parent
	}
	return isChildOfForm(n.Parent)
}
