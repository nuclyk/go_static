package main

import "testing"

func TestSplitNodesDelimiterBold(t *testing.T) {
	node := textNode{
		text:     "Some **bold** text.",
		textType: TEXT,
		url:      "",
	}

	nodes := []textNode{
		{text: "Some ", textType: TEXT},
		{text: "bold", textType: BOLD},
		{text: " text.", textType: TEXT},
	}

	testNodes, err := splitNodesDelimiter([]textNode{node}, "**", BOLD)

	if err != nil {
		t.Errorf("Error when converting: %v", err)
	}

	for i, node := range nodes {
		if node != testNodes[i] {
			t.Errorf("Expected text nodes are wrong\nExpected: %v\nActual: %v", nodes, testNodes)
		}
	}
}

func TestSplitNodesDelimiterItalic(t *testing.T) {
	node := textNode{
		text:     "Some _italic_ text.",
		textType: TEXT,
		url:      "",
	}

	nodes := []textNode{
		{text: "Some ", textType: TEXT},
		{text: "italic", textType: ITALIC},
		{text: " text.", textType: TEXT},
	}

	testNodes, err := splitNodesDelimiter([]textNode{node}, "_", ITALIC)

	if err != nil {
		t.Errorf("Error when converting: %v", err)
	}

	for i, node := range nodes {
		if node != testNodes[i] {
			t.Errorf("Expected text nodes are wrong\nExpected: %v\nActual: %v", nodes, testNodes)
		}
	}
}

func TestSplitNodesDelimiterCode(t *testing.T) {
	node := textNode{
		text:     "Some `code` text.",
		textType: TEXT,
		url:      "",
	}

	nodes := []textNode{
		{text: "Some ", textType: TEXT},
		{text: "code", textType: CODE},
		{text: " text.", textType: TEXT},
	}

	testNodes, err := splitNodesDelimiter([]textNode{node}, "`", CODE)

	if err != nil {
		t.Errorf("Error when converting: %v", err)
	}

	for i, node := range nodes {
		if node != testNodes[i] {
			t.Errorf("Expected text nodes are wrong\nExpected: %v\nActual: %v", nodes, testNodes)
		}
	}
}

func TestSplitNodesDelimiterCombined(t *testing.T) {
	node := textNode{
		text:     "Some **bold** text with some _italics_ and additional `code`.",
		textType: TEXT,
		url:      "",
	}

	nodes := []textNode{
		{text: "Some ", textType: TEXT},
		{text: "bold", textType: BOLD},
		{text: " text with some ", textType: TEXT},
		{text: "italics", textType: ITALIC},
		{text: " and additional ", textType: TEXT},
		{text: "code", textType: CODE},
		{text: ".", textType: TEXT},
	}

	testNodes, err := splitNodesDelimiter([]textNode{node}, "**", BOLD)
	testNodes, err = splitNodesDelimiter(testNodes, "_", ITALIC)
	testNodes, err = splitNodesDelimiter(testNodes, "`", CODE)

	if err != nil {
		t.Errorf("Error when converting: %v", err)
	}

	for i, node := range nodes {
		if node != testNodes[i] {
			t.Errorf("Expected text nodes are wrong\nExpected: %v\nActual: %v", nodes, testNodes)
		}
	}
}
