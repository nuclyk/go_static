package main

import (
	"errors"
	"fmt"
)

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

func textNodeToHtmlNode(tn textNode) (leafNode, error) {
	if tn.textType == TEXT {
		return leafNode{tag: "", value: tn.text, attributes: nil}, nil
	} else if tn.textType == BOLD {
		return leafNode{tag: "b", value: tn.text, attributes: nil}, nil
	} else if tn.textType == ITALIC {
		return leafNode{tag: "i", value: tn.text, attributes: nil}, nil
	} else if tn.textType == CODE {
		return leafNode{tag: "code", value: tn.text, attributes: nil}, nil
	} else if tn.textType == LINK {
		return leafNode{tag: "a", value: tn.text, attributes: map[string]string{"href": tn.url}}, nil
	} else if tn.textType == IMAGE {
		return leafNode{tag: "img", value: "", attributes: map[string]string{"src": tn.url, "alt": tn.text}}, nil
	}

	return leafNode{}, errors.New("invalid text type")
}
