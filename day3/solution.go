package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parseInput() string {
	dataBytes, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(dataBytes)
}

func solve_challenge_1(data string) {
	matched, _ := matchString(data)
	
	sum := 0

	// iterate over the matched 2d matrix
	for _, row := range matched {
		// for each column multiply and add to sum
			// parse second and third elems to ints
			x, err := strconv.Atoi(row[1]);
			y, err := strconv.Atoi(row[2]);
				
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			sum += (x*y);
	}

	fmt.Println("sum:", sum)
}

func solve_challenge_2(data string) {
	

	matches, re_dont, re_do, re_mul := matchWithCond(data)
	is_enabled := true
	sum := 0

	for _, match := range matches {
		instruction := match[0]

		// Handle `do()` and `don't()` instructions
		if re_do.MatchString(instruction) {
			is_enabled = true
		} else if re_dont.MatchString(instruction) {
			is_enabled = false
		} else if mulMatch := re_mul.FindStringSubmatch(instruction); mulMatch != nil {
			// Handle `mul(X, Y)`
			if is_enabled {
				x, _ := strconv.Atoi(mulMatch[1])
				y, _ := strconv.Atoi(mulMatch[2])
				sum += x * y
			}
		}
	}

	fmt.Println("sum", sum)
}

func matchWithCond(text string) ([][]string, *regexp.Regexp, *regexp.Regexp, *regexp.Regexp) {
	pattern_dont := `don't\(\)`
	pattern_do := `do\(\)`
	pattern_mul := `mul\((\d{1,3}),(\d{1,3})\)`
	
	re_dont := regexp.MustCompile(pattern_dont)
	re_do := regexp.MustCompile(pattern_do)
	re_mul := regexp.MustCompile(pattern_mul)

	matches := regexp.MustCompile(pattern_mul + "|" + pattern_do + "|" + pattern_dont).FindAllStringSubmatch(text, -1)	

	return matches, re_dont, re_do, re_mul

}

func matchString(text string) ([][]string, [][]int) {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)

	matched := re.FindAllStringSubmatch(text, -1)
	matchIdx := re.FindAllStringSubmatchIndex(text, -1)

	return matched, matchIdx
}

func main() {
	data := parseInput()
	// solve_challenge_1(data)
	solve_challenge_2(data)
}
