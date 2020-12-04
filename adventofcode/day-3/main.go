package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	rightSteps := 3
	globalPosition := 0
	slopeMap := parseSlopeMap()
	slopeMapWidth := len(slopeMap[0])

	hits := 0

	for _, slopeRow := range slopeMap[1:] {
		globalPosition = globalPosition + rightSteps
		localPosition := globalPosition % slopeMapWidth
		plotValue := []rune(slopeRow)[localPosition]

		if plotValue == '#' {
			hits = hits + 1
		}
	}

	fmt.Printf("%v Hits", hits)
}

func parseSlopeMap() []string {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var slopeMap []string

	for scanner.Scan() {
		slopeMap = append(slopeMap, scanner.Text())
	}

	return slopeMap
}
