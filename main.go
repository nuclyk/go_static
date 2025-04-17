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
		text:     "Some **bolded** text with some _italics_ and `code`.",
		textType: TEXT,
		url:      "",
	}

	// t3 := textNode{
	// 	text:     "Just a normal text",
	// 	textType: TEXT,
	// 	url:      "",
	// }

	fmt.Println(t1.equals(t2))
	fmt.Println(html_node)
	fmt.Println(parent_node)
	fmt.Println(leaf_node)

	t1html, err := textNodeToHtmlNode(t1)
	if err == nil {
		fmt.Println(t1html)
	}

	testNodes := []textNode{}
	testNodes = append(testNodes, t2)
	testNodes, err = splitNodesDelimiter(testNodes, "**", BOLD)
	testNodes, err = splitNodesDelimiter(testNodes, "_", ITALIC)
	testNodes, err = splitNodesDelimiter(testNodes, "`", CODE)
	if err == nil {
		fmt.Println("FINAL VERSION")
		fmt.Println(testNodes)
	}

	// image := "This is text with an ![image](https://i.imgur.com/zjjcJKZ.png)." +
	// 	"And this is another one: ![image](https://i.imgur.com/zjjcJKZ.png)"
	// result := extractMarkdownImages(image)
	// fmt.Print(result)

	link := string("This is text with a [link](https://boot.dev) and [another link](https://blog.boot.dev)")
	links := extractMarkdownLinks(link)
	fmt.Println(links)

}
