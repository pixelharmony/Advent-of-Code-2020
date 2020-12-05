package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func main() {
	passports := make(chan map[string]string)

	go start(passports)

	validCount := 0

	for range passports {
		validCount++
	}

	fmt.Printf("Found %v valid passports", validCount)
}

// start reads passport data from file and trigger async processing
func start(passports chan map[string]string) {

	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(scanBlankLines)

	var wg sync.WaitGroup

	for scanner.Scan() {
		scannedText := scanner.Text()

		wg.Add(1)
		go processPassporttext(scannedText, passports, &wg)
	}

	wg.Wait()
	close(passports)
}

// processPassportText converts the raw text into a map of passport values
func processPassporttext(text string, channel chan map[string]string, wg *sync.WaitGroup) {
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
	validators := []validator{
		requiredFieldsValidator,
		minMaxValidator(1920, 2002, "byr"),
		minMaxValidator(2010, 2020, "iyr"),
		minMaxValidator(2020, 2030, "eyr"),
		hgtValidator,
		regexValidator(`^#[A-f0-9]{6}$`, "hcl"),
		regexValidator(`^amb|blu|brn|gry|grn|hzl|oth$`, "ecl"),
		regexValidator(`^\d{9}$`, "pid")}

	for _, validator := range validators {
		if !validator(passport) {
			return false
		}
	}

	return true
}

// passport validator functions

type validator func(passport map[string]string) bool

func requiredFieldsValidator(passport map[string]string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range requiredFields {
		if _, ok := passport[field]; !ok {
			return false
		}
	}

	return true
}

func minMaxValidator(min int, max int, key string) func(map[string]string) bool {
	return func(passport map[string]string) bool {
		value, _ := strconv.Atoi(passport[key])

		if value >= min && value <= max {
			return true
		}

		return false
	}
}

func hgtValidator(passport map[string]string) bool {
	rp := regexp.MustCompile(`(\d+)(cm|in)`)

	if !rp.MatchString(passport["hgt"]) {
		return false
	}

	match := rp.FindStringSubmatch(passport["hgt"])
	value, _ := strconv.Atoi(match[1])
	unit := match[2]

	if unit == "cm" && value >= 150 && value <= 193 {
		return true
	}

	if unit == "in" && value >= 59 && value <= 76 {
		return true
	}

	return false
}

func regexValidator(pattern string, key string) func(map[string]string) bool {
	return func(passport map[string]string) bool {
		rp := regexp.MustCompile(pattern)

		if !rp.MatchString(passport[key]) {
			return false
		}

		return true
	}
}

// scanBlankLines splits the scanner input when a line is empty
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
