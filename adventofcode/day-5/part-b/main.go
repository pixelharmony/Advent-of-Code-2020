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
	grid := createGrid(rows, seats)

	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	foundSeatIDs := make(map[int]bool)

	for scanner.Scan() {

		boardingPass := scanner.Text()

		rowCodes, seatCodes := parseBoardingPass(boardingPass)

		row := axisBinarySearch(rows, rowCodes)
		seat := axisBinarySearch(seats, seatCodes)
		seatID := seatID(row, seat)
		foundSeatIDs[seatID] = true

		grid[row][seat] = true
	}

	emptySeatIDs := findEmptySeats(grid)

	for _, i := range emptySeatIDs {
		if prevSeat, nextSeat := foundSeatIDs[i-1], foundSeatIDs[i+1]; prevSeat && nextSeat {
			fmt.Println(i)
		}
	}
}

func findEmptySeats(grid map[int]map[int]bool) []int {

	var unfilledSeatIDs []int

	for r := range grid {
		for s, filled := range grid[r] {
			if !filled {
				unfilledSeatID := seatID(r, s)
				unfilledSeatIDs = append(unfilledSeatIDs, unfilledSeatID)
			}
		}
	}

	return unfilledSeatIDs
}

func createGrid(rows []int, seats []int) map[int]map[int]bool {

	grid := make(map[int]map[int]bool)

	for _, r := range rows {

		grid[r] = make(map[int]bool)

		for _, s := range seats {
			grid[r][s] = false
		}
	}

	return grid
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
