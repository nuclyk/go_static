package main

import (
	"errors"
	"regexp"
	"slices"
	"strings"
)

func splitNodesDelimiter(oldNodes []textNode, delimiter string, textType textType) ([]textNode, error) {
	newNodes := []textNode{}
	for _, oldNode := range oldNodes {
		if oldNode.textType != TEXT {
			newNodes = append(newNodes, oldNode)
			continue
		}

		var splitNodes []textNode
		sections := strings.Split(oldNode.text, delimiter)

		if len(sections)%2 == 0 {
			return nil, errors.New("invalid markdown, formatted section not closed")
		}

		for i := range len(sections) {
			if sections[i] == "" {
				continue
			}
			if i%2 == 0 {
				node := textNode{text: sections[i], textType: TEXT, url: ""}
				splitNodes = append(splitNodes, node)
			} else {
				node := textNode{text: sections[i], textType: textType, url: ""}
				splitNodes = append(splitNodes, node)
			}
		}
		newNodes = slices.Concat(newNodes, splitNodes)
	}
	return newNodes, nil
}

// func spiltNodesImage(oldNodes []textNode) {
// 	newNodes = []textNode{}
// 	for _, oldNode := range oldNodes {
// 		if oldNode.textType != TEXT {
// 			newNodes = append(newNodes, oldNode)
// 			continue
// 		}
// 	}
// }

func extractMarkdownImages(text string) [][]string {
	re := regexp.MustCompile(`!\[([^\[\]]*)\]\(([^\(\)]*)\)`)
	found := re.FindAllStringSubmatch(text, -1)
	var filtered [][]string
	for _, img := range found {
		filtered = append(filtered, []string{img[1], img[2]})
	}
	return filtered
}

// Go doesn't support lookbehind in regex
func extractMarkdownLinks(text string) [][]string {
	re := regexp.MustCompile(`[^!]\[(.*?)\]\((.*?)\)`)

	found := re.FindAllStringSubmatch(text, -1)
	var filtered [][]string
	for _, img := range found {
		filtered = append(filtered, []string{img[1], img[2]})
	}
	return filtered
}
