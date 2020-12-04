package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type values struct {
	A int
	B int
}

func main() {

	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var prevValues = make(map[int]int)
	var computedValues = make(map[int]values)

	for scanner.Scan() {
		var value, _ = strconv.Atoi(scanner.Text())

		var matchingValue = 2020 - value

		if computedValue, ok := computedValues[matchingValue]; ok {
			fmt.Printf("Match Found: %v %v %v, Answer: %v\n", value, computedValue.A, computedValue.B, value*computedValue.A*computedValue.B)
		}

		for _, prevValue := range prevValues {
			var result = value + prevValue
			computedValues[result] = values{prevValue, value}
		}

		prevValues[value] = value
	}
}
