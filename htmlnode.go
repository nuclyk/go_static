package main

import (
	"errors"
	"fmt"
)

type node interface {
	attributestoHTML() string
	toHTML() (string, error)
}

// htmlNode can have children
// Implements Node interface
type htmlNode struct {
	tag        string
	value      string
	attributes map[string]string
	children   []node
}

func (hm htmlNode) attributestoHTML() string {
	if len(hm.attributes) == 0 {
		return ""
	}

	attr_html := ""
	for k, v := range hm.attributes {
		attr_html += fmt.Sprintf(" %s=\"%s\"", k, v)
	}
	return attr_html
}

func (hm htmlNode) toHTML() (string, error) {
	return "", errors.New("toHTML() ot implemented in htmlNode.")
}

func (hm htmlNode) String() string {
	attrs := ""
	for k, v := range hm.attributes {
		attrs += fmt.Sprintf("%s:%s, ", k, v)
	}
	return fmt.Sprintf("htmlNode(%s, %s%s)", hm.tag, attrs, hm.value)
}

// leafNode, which means it has no children.
// Implements Node interface.
type leafNode struct {
	tag        string
	value      string
	attributes map[string]string
}

func (ln leafNode) attributestoHTML() string {
	if len(ln.attributes) == 0 {
		return ""
	}

	attr_html := ""
	for k, v := range ln.attributes {
		attr_html += fmt.Sprintf("%s=\"%s\"", k, v)
	}
	return attr_html
}

// leafNode method that converts leafNode into HTML string.
// Returns error if tag or value is not set.
func (ln leafNode) toHTML() (string, error) {
	if len(ln.tag) == 0 {
		return "", errors.New("leafNode: Invalid HTML: no tag")
	}
	if len(ln.value) == 0 {
		return "", errors.New("leafNode: Invalid HTML: no value")
	}
	return fmt.Sprintf("<%s %s>%s</%s>", ln.tag, ln.attributestoHTML(), ln.value, ln.tag), nil
}

func (ln leafNode) String() string {
	attrs := ""
	for k, v := range ln.attributes {
		attrs += fmt.Sprintf("%s:%s, ", k, v)
	}
	return fmt.Sprintf("leafNode(%s, %s%s)", ln.tag, attrs, ln.value)
}

// parentNode
// Implements Node interface
type parentNode struct {
	tag        string
	attributes map[string]string
	children   []node
}

func (pn parentNode) attributestoHTML() string {
	if len(pn.attributes) == 0 {
		return ""
	}

	attr_html := ""
	for k, v := range pn.attributes {
		attr_html += fmt.Sprintf("%s=\"%s\"", k, v)
	}
	return attr_html
}

func (pn parentNode) toHTML() (string, error) {
	if len(pn.tag) == 0 {
		return "", errors.New("parentNode: Invalid HTML: no tag")
	}
	if len(pn.children) == 0 {
		return "", errors.New("parentNode: Invalid HTML: no children")
	}

	children_html := ""
	for _, child := range pn.children {
		html, err := child.toHTML()
		if err == nil {
			children_html += html
		}
	}
	return fmt.Sprintf("<%s %s>%s</%s>", pn.tag, pn.attributestoHTML(), children_html, pn.tag), nil
}

func (pn parentNode) String() string {
	return fmt.Sprintf("parentNode(%s, children: %s, %s)", pn.tag, pn.children, pn.attributes)
}
