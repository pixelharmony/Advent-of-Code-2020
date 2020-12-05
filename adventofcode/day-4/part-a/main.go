package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	passports := make(chan map[string]string)

	go start(passports)

	validCount := 0

	for i := range passports {
		fmt.Println("Valid Passport")
		fmt.Println(i)
		validCount++
	}

	fmt.Printf("Found %v valid passports", validCount)
}

func start(passports chan map[string]string) {

	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(scanBlankLines)

	var wg sync.WaitGroup

	for scanner.Scan() {
		scannedText := scanner.Text()

		wg.Add(1)
		go processPasswordtext(scannedText, passports, &wg)
	}

	wg.Wait()
	close(passports)
}

func processPasswordtext(text string, channel chan map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()
	spacesOnlyText := strings.ReplaceAll(text, "\n", " ")
	passportProps := strings.Split(spacesOnlyText, " ")

	passport := make(map[string]string)

	for _, stringProp := range passportProps {
		prop := strings.Split(stringProp, ":")
		passport[prop[0]] = prop[1]
	}

	if isValidPassport(passport) {
		channel <- passport
	}

}

func isValidPassport(passport map[string]string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range requiredFields {
		if _, ok := passport[field]; !ok {
			return false
		}
	}

	return true
}

func scanBlankLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	searchBytes := []byte("\n\n")

	if i := bytes.Index(data, searchBytes); i >= 0 {
		return i + 2, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}
