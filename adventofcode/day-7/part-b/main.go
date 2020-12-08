package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type node struct {
	Parents  []string
	Children []bag
}

type bag struct {
	Count  int
	Colour string
}

func main() {
	graph := buildGraph()
	shinyGoldNode := getNode(graph, "shiny gold")
	count := countBags(graph, shinyGoldNode)

	fmt.Println(count)
}

func countBags(graph map[string]*node, node *node) int {
	bagCount := 0

	for _, child := range node.Children {
		bagCount = bagCount + child.Count

		childNode := getNode(graph, child.Colour)
		innerBags := countBags(graph, childNode)

		bagCount = bagCount + (innerBags * child.Count)
	}

	return bagCount
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
			childNode := getNode(graph, child.Colour)
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

func parseRule(text string) (parent string, children []bag) {
	rp := regexp.MustCompile(`((\d)*(?:\s)*(\w+ \w+) bag)`)
	match := rp.FindAllStringSubmatch(text, -1)

	parent = match[0][3]

	for i := 1; i < len(match); i++ {
		colour := match[i][3]
		if colour != "no other" {
			count, _ := strconv.Atoi(match[i][2])
			children = append(children, bag{count, colour})
		}
	}

	return parent, children
}
