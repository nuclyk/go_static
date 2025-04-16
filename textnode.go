package main

import "fmt"

type Equaler interface {
	equals(node textNode) bool
}

type textType string

const (
	TEXT   textType = "text"
	BOLD   textType = "bold"
	ITALIC textType = "italic"
	CODE   textType = "code"
	LINK   textType = "link"
	IMAGE  textType = "image"
)

type textNode struct {
	text     string
	textType textType
	url      string
}

func (tn textNode) String() string {
	return fmt.Sprintf("TextNode(%s, %s, %s)", tn.text, tn.textType, tn.url)
}

func (tn textNode) equals(other textNode) bool {
	return tn.text == other.text && tn.textType == other.textType && tn.url == other.url
}
