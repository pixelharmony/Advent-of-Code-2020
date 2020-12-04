package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passwordStruct struct {
	PolicyMin int
	PolicyMax int
	Character string
	Password  string
}

var rp = regexp.MustCompile(`(\d+)-(\d+) (\w): (.+)`)

func main() {

	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var validPasswordCount = 0

	for scanner.Scan() {

		var value = parsePassword(scanner.Text())

		var letterCount = strings.Count(value.Password, value.Character)

		if letterCount >= value.PolicyMin && letterCount <= value.PolicyMax {
			validPasswordCount = validPasswordCount + 1
		}
	}

	fmt.Printf("%v valid passwords", validPasswordCount)
}

func parsePassword(value string) passwordStruct {
	var match = rp.FindAllStringSubmatch(value, 3)[0]

	var policyMin, _ = strconv.Atoi(match[1])
	var policyMax, _ = strconv.Atoi(match[2])
	var character = match[3]
	var password = match[4]

	return passwordStruct{policyMin, policyMax, character, password}
}
