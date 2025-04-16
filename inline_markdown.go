package main

import (
	"errors"
	"fmt"
	"strings"
)

func splitNodesDelimiter(oldNodes []textNode, delimiter string, textType textType) ([]textNode, error) {
	newNodes := []textNode{}
	for _, oldNode := range oldNodes {
		// fmt.Println(oldNode)
		if oldNode.textType != TEXT {
			newNodes = append(newNodes, oldNode)
			continue
		}

		var splitNodes []textNode
		sections := strings.Split(oldNode.text, delimiter)
		fmt.Printf("Sections: %v\n", sections)

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
		fmt.Printf("Split nodes: %v\n", splitNodes)
		for _, node := range splitNodes {
			newNodes = append(newNodes, node)
		}
	}
	fmt.Println(newNodes)

	return newNodes, nil
}
