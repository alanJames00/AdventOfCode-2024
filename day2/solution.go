package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// remove an item from the slice
func removeItem(slice []int, idx int) []int {
	// bound check
	if 0 <= idx && idx < len(slice) {
		// slice the slice
		// create a copy slice
		copySlice := make([]int, 0, len(slice)-1)
		copySlice = append(copySlice, slice[:idx]...)
		copySlice = append(copySlice, slice[idx+1:]...)

		return copySlice
	} else {
		fmt.Printf("slice out of bounds with slice of len %v and idx: %v\n", len(slice), idx)
		os.Exit(1)
		return nil
	}
}

// parse input
func parseInput() [][]int {
	// read from file
	dataBytes, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("error parsing input", err)
	}

	var result [][]int

	// parse dataBytes to string slice
	lines := strings.Split(string(dataBytes), "\n")

	for i, line := range lines {
		// skip i = 1000
		if i == 1000 {
			continue
		}
		// parse line into string slice
		strSlice := strings.Fields(line)
		// parse strSlice to int slice
		var temp []int
		for _, val := range strSlice {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("error parsing integer", err)
				os.Exit(1)
			}
			temp = append(temp, intVal)
		}
		// push the temp arr to result 2d slice
		result = append(result, temp)
	}

	return result
}

func solve_challenge_1(data [][]int) {

	unsafeCount := 0

	// iterate over each report

	for _, row := range data {
		isIncreasing := false

		lastVal := row[0]

		// iterate over each values of row
		unsafe := false
		for j, val := range row {
			if j == 1 {
				if lastVal > val {
					// decreasing
					isIncreasing = false
				} else {
					isIncreasing = true
				}

				// check for allowed diff
				diff := math.Abs(float64(lastVal) - float64(val))
				if diff < 1 || diff > 3 {
					unsafe = true
				}
				lastVal = val
			} else if j > 1 {
				// has to be following same monotonic trend
				if isIncreasing && lastVal > val || !isIncreasing && lastVal < val {
					// unsafe
					unsafe = true
				}

				diff := math.Abs(float64(lastVal) - float64(val))
				if diff < 1 || diff > 3 {
					unsafe = true
				}

				lastVal = val
			}

		}

		// check dampener results
	
		if unsafe && !dampener(row) {
			unsafeCount++
		}
		fmt.Println(row, unsafe)
	}

	fmt.Println(len(data) - unsafeCount)
}

func dampener(Ogrow []int) bool {
	// iterate and remove one item from the row at a time
	for i := range Ogrow {
		row := removeItem(Ogrow, i)

		// safety check
		isIncreasing := false

		lastVal := row[0]

		// iterate over each values of row
		unsafe := false
		for j, val := range row {
			if j == 1 {
				if lastVal > val {
					// decreasing
					isIncreasing = false
				} else {
					isIncreasing = true
				}

				// check for allowed diff
				diff := math.Abs(float64(lastVal) - float64(val))
				if diff < 1 || diff > 3 {
					unsafe = true
				}
				lastVal = val
			} else if j > 1 {
				// has to be following same monotonic trend
				if isIncreasing && lastVal > val || !isIncreasing && lastVal < val {
					// unsafe
					unsafe = true
				}

				diff := math.Abs(float64(lastVal) - float64(val))
				if diff < 1 || diff > 3 {
					unsafe = true
				}

				lastVal = val
			}

		}
		fmt.Println(row, unsafe)
		// return true if one of them is safe
		if(!unsafe) {
			return true
		}
	}

	// return false
	return false
}
func main() {
	data := parseInput()
	solve_challenge_1(data)
}
