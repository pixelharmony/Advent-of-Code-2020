package main

import (
	"bufio"
	"fmt"
	"os"
)

var globalPosition = 0

func main() {

	hits := 1

	slopeMap := parseSlopeMap()

	hits = hits * calcSlopeHits(slopeMap, 1, 1)
	hits = hits * calcSlopeHits(slopeMap, 3, 1)
	hits = hits * calcSlopeHits(slopeMap, 5, 1)
	hits = hits * calcSlopeHits(slopeMap, 7, 1)
	hits = hits * calcSlopeHits(slopeMap, 1, 2)

	fmt.Printf("Hit %v trees", hits)
}

func calcSlopeHits(slopeMap []string, rightSteps int, downSteps int) int {
	slopeMapWidth := len(slopeMap[0])

	hits := 0

	for i := downSteps; i < len(slopeMap); i += downSteps {
		slopeRow := slopeMap[i]
		globalPosition = globalPosition + rightSteps
		localPosition := globalPosition % slopeMapWidth
		plotValue := []rune(slopeRow)[localPosition]

		if plotValue == '#' {
			hits = hits + 1
		}
	}

	globalPosition = 0
	return hits
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
