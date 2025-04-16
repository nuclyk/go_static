package main

import "fmt"

func main() {
	html_node := htmlNode{
		tag:   "div",
		value: "Some sort of text.",
		attributes: map[string]string{
			"class": "bg-color",
		},
		children: []node{},
	}

	leaf_node := leafNode{
		tag:   "a",
		value: "Link to google",
		attributes: map[string]string{
			"href":  "https://google.com",
			"class": "beautify",
		},
	}

	parent_node := parentNode{
		tag:        "div",
		attributes: nil,
		children:   []node{leaf_node},
	}

	t1 := textNode{
		text:     "Linkt to google",
		textType: LINK,
		url:      "http://google.com",
	}

	t2 := textNode{
		text:     "Some text",
		textType: BOLD,
		url:      "http://google.com",
	}

	fmt.Println(t1.equals(t2))
	fmt.Println(html_node)
	fmt.Println(parent_node)
	fmt.Println(leaf_node)

	t1html, err := textNodeToHtmlNode(t1)
	if err == nil {
		fmt.Println(t1html)
	}

}
