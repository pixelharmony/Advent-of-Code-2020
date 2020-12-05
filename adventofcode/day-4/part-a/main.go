package main

import (
	"bufio"
	"os"
)

func main() {

}

func readInput() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
	}
}
