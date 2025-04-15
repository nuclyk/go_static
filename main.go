package main

import "fmt"

func main() {
	html_node := HTMLNode{
		tag:   "div",
		value: "Some sort of text.",
		attributes: map[string]string{
			"class": "bg-color",
		},
		children: []Node{},
	}

	leaf_node := LeafNode{
		tag:   "a",
		value: "Link to google",
		attributes: map[string]string{
			"href": "https://google.com",
		},
	}

	parent_node := ParentNode{
		tag:        "div",
		attributes: nil,
		children:   []Node{leaf_node},
	}

	fmt.Println(html_node)
	fmt.Println(leaf_node.ToHTML())
	fmt.Println(parent_node.ToHTML())

}
