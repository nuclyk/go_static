package main

import (
	"errors"
	"fmt"
)

type Node interface {
	AttributesToHTML() string
	ToHTML() (string, error)
}

// HTMLNode can have children
// Implements Node interface
type HTMLNode struct {
	tag        string
	value      string
	attributes map[string]string
	children   []Node
}

func (hm HTMLNode) AttributesToHTML() string {
	if len(hm.attributes) == 0 {
		return ""
	}

	attr_html := ""
	for k, v := range hm.attributes {
		attr_html += fmt.Sprintf(" %s=\"%s\"", k, v)
	}
	return attr_html
}

func (hm HTMLNode) ToHTML() (string, error) {
	return "", errors.New("ToHTML() ot implemented in HTMLNode.")
}

func (hm HTMLNode) String() string {
	return fmt.Sprintf("HTMLNode(%s, %s, %s)", hm.tag, hm.attributes, hm.value)
}

// LeafNode, which means it has no children.
// Implements Node interface.
type LeafNode struct {
	tag        string
	value      string
	attributes map[string]string
}

func (ln LeafNode) AttributesToHTML() string {
	if len(ln.attributes) == 0 {
		return ""
	}

	attr_html := ""
	for k, v := range ln.attributes {
		attr_html += fmt.Sprintf("%s=\"%s\"", k, v)
	}
	return attr_html
}

// LeafNode method that converts LeafNode into HTML string.
// Returns error if tag or value is not set.
func (ln LeafNode) ToHTML() (string, error) {
	if len(ln.tag) == 0 {
		return "", errors.New("LeafNode: Invalid HTML: no tag")
	}
	if len(ln.value) == 0 {
		return "", errors.New("LeafNode: Invalid HTML: no value")
	}
	return fmt.Sprintf("<%s %s>%s</%s>", ln.tag, ln.AttributesToHTML(), ln.value, ln.tag), nil
}

func (ln LeafNode) String() string {
	return fmt.Sprintf("LeafNode(%s, %s, %s)", ln.tag, ln.attributes, ln.value)
}

// ParentNode
// Implements Node interface
type ParentNode struct {
	tag        string
	attributes map[string]string
	children   []Node
}

func (pn ParentNode) AttributesToHTML() string {
	if len(pn.attributes) == 0 {
		return ""
	}

	attr_html := ""
	for k, v := range pn.attributes {
		attr_html += fmt.Sprintf("%s=\"%s\"", k, v)
	}
	return attr_html
}

func (pn ParentNode) ToHTML() (string, error) {
	if len(pn.tag) == 0 {
		return "", errors.New("ParentNode: Invalid HTML: no tag")
	}
	if len(pn.children) == 0 {
		return "", errors.New("ParentNode: Invalid HTML: no children")
	}

	children_html := ""
	for _, child := range pn.children {
		html, err := child.ToHTML()
		if err == nil {
			children_html += html
		}
	}
	return fmt.Sprintf("<%s %s>%s</%s>", pn.tag, pn.AttributesToHTML(), children_html, pn.tag), nil
}

func (pn ParentNode) String() string {
	return fmt.Sprintf("ParentNode(%s, children: %s, %s)", pn.tag, pn.children, pn.attributes)
}
