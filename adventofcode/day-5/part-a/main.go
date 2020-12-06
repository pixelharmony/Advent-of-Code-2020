package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	keepLower = "KeepLower"
	keepUpper = "KeepUpper"
)

func main() {

	rowCount := 128
	seatCount := 8
	rows := createAxis(rowCount)
	seats := createAxis(seatCount)

	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	highestSeatID := 0

	for scanner.Scan() {

		boardingPass := scanner.Text()

		rowCodes, seatCodes := parseBoardingPass(boardingPass)

		row := axisBinarySearch(rows, rowCodes)
		seat := axisBinarySearch(seats, seatCodes)
		seatID := seatID(row, seat)

		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	fmt.Printf("Highest Seat ID: %v\n", highestSeatID)
}

func createAxis(count int) []int {
	axis := make([]int, count)

	for i := 0; i < count; i++ {
		axis[i] = i
	}

	return axis
}

func parseBoardingPass(boardingPass string) (rowCodes []rune, seatCodes []rune) {
	runes := []rune(boardingPass)
	rowCodes = runes[0:7]
	seatCodes = runes[7:10]

	return rowCodes, seatCodes
}

func axisBinarySearch(axis []int, boardingPassCodes []rune) int {
	axisRemainder := axis

	for _, bpc := range boardingPassCodes {
		bound := codeToBound(bpc)
		axisRemainder = bisect(axisRemainder, bound)
	}

	return axisRemainder[0]
}

func codeToBound(rowCode rune) string {
	if rowCode == 'F' || rowCode == 'L' {
		return keepLower
	}
	return keepUpper
}

func bisect(array []int, keepBound string) []int {

	sliceSize := len(array) / 2

	if keepBound == keepLower {
		return array[:sliceSize]
	}
	return array[sliceSize:]
}

func seatID(row int, seat int) int {
	return row*8 + seat
}
