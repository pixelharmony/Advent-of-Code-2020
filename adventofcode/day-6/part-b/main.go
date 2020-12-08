package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(scanBlankLines)

	total := 0

	for scanner.Scan() {
		text := scanner.Text()
		groupSize := countGroups(text)
		answers := []rune(clean(text))
		groupedAnswers := groupRunes(answers)
		answerCount := countAnswers(groupSize, groupedAnswers)

		total = total + answerCount
	}

	fmt.Printf("Total: %v", total)
}

func countAnswers(groupSize int, groupedAnswers map[rune][]rune) int {
	answers := 0

	for _, i := range groupedAnswers {
		if len(i) == groupSize {
			answers++
		}
	}

	return answers
}

func groupRunes(runes []rune) map[rune][]rune {
	groupedRunes := make(map[rune][]rune)

	for _, i := range runes {
		groupedRunes[i] = append(groupedRunes[i], i)
	}

	return groupedRunes
}

func clean(text string) string {
	withoutNewLines := strings.ReplaceAll(text, "\n", "")
	withoutSpaces := strings.ReplaceAll(withoutNewLines, " ", "")
	return withoutSpaces
}

func countGroups(text string) int {
	return len(strings.Split(text, "\n"))
}
