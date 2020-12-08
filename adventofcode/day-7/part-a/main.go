package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type node struct {
	Parents  []string
	Children []string
}

func main() {
	graph := buildGraph()

	var nodesToProcess []*node

	shinyGoldNode := getNode(graph, "shiny gold")
	nodesToProcess = append(nodesToProcess, shinyGoldNode)
	foundNodes := make(map[string]struct{})

	for len(nodesToProcess) > 0 {
		node := pop(&nodesToProcess)

		for _, parent := range node.Parents {
			parentNode := getNode(graph, parent)
			nodesToProcess = append(nodesToProcess, parentNode)

			foundNodes[parent] = struct{}{}
		}
	}

	fmt.Printf("Total: %v", len(foundNodes))
}

func pop(nodesToProcess *[]*node) *node {
	length := len(*nodesToProcess)
	popped := (*nodesToProcess)[length-1]
	*nodesToProcess = append((*nodesToProcess)[:length-1])
	return popped
}

func buildGraph() map[string]*node {
	content, _ := ioutil.ReadFile("./input.txt")
	rules := strings.Split(string(content), "\n")

	graph := make(map[string]*node)

	for _, rule := range rules {
		parent, children := parseRule(rule)
		parentNode := getNode(graph, parent)
		parentNode.Children = append(parentNode.Children, children...)

		for _, child := range children {
			childNode := getNode(graph, child)
			childNode.Parents = append(childNode.Parents, parent)
		}
	}

	return graph
}

func getNode(graph map[string]*node, key string) *node {
	if _, ok := graph[key]; !ok {
		graph[key] = &node{}
	}

	return graph[key]
}

func parseRule(text string) (parent string, children []string) {
	rp := regexp.MustCompile(`((\d)*(?:\s)*(\w+ \w+) bag)`)
	match := rp.FindAllStringSubmatch(text, -1)

	parent = match[0][3]

	for i := 1; i < len(match); i++ {
		colour := match[i][3]
		children = append(children, colour)
	}

	return parent, children
}
