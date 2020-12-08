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
		cleanText := clean(text)
		count := countAnswers(cleanText)

		total = total + count
	}

	fmt.Printf("Total: %v", total)
}

func countAnswers(text string) int {
	answers := make(map[rune]struct{})

	for _, char := range text {
		answers[char] = struct{}{}
	}

	return len(answers)
}

func clean(text string) string {
	return strings.ReplaceAll(text, "\n", "")
}
