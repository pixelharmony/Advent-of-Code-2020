package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type passwordStruct struct {
	PolicyMin int
	PolicyMax int
	Character rune
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

		var characters = []rune(value.Password)

		var validPolicyMin = characters[value.PolicyMin-1] == value.Character
		var validPolicyMax = characters[value.PolicyMax-1] == value.Character

		if validPolicyMin != validPolicyMax {
			validPasswordCount = validPasswordCount + 1
		}
	}

	fmt.Printf("%v valid passwords", validPasswordCount)
}

func parsePassword(value string) passwordStruct {
	var match = rp.FindAllStringSubmatch(value, 3)[0]

	var policyMin, _ = strconv.Atoi(match[1])
	var policyMax, _ = strconv.Atoi(match[2])
	var character = rune(match[3][0])
	var password = match[4]

	return passwordStruct{policyMin, policyMax, character, password}
}
