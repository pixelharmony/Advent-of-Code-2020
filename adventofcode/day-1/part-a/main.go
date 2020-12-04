package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var prevValues = make(map[int]bool)

	for scanner.Scan() {
		var value, _ = strconv.Atoi(scanner.Text())

		var matchingValue = 2020 - value

		if _, ok := prevValues[matchingValue]; ok {
			fmt.Println(value * matchingValue)
		}

		prevValues[value] = true
	}
}
