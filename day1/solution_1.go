package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile() (string, error) {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func parseInput(data string) ([]int, []int) {	
		lines := strings.Split(data, "\n")

	// keep track of left and right lists
	var leftList []int
	var rightList []int

	for i, line := range lines {
		// skip i = 1000
		if i == 1000 {
			continue
		}
		lInt, err := strconv.Atoi(strings.Fields(line)[0])
		rInt, err := strconv.Atoi(strings.Fields(line)[1])
		if err != nil {
			fmt.Println("error parsing int")
			os.Exit(1)
		}
		leftList = append(leftList, lInt)
		rightList = append(rightList, rInt)
	}

	// sort left and right lists
	sort.Ints(leftList)
	sort.Ints(rightList)

	return leftList, rightList;

}

func solve_challange_1(leftList []int, rightList []int) {

	// iterate over the both lists and adding up the distance
	distance := 0
	for i := 0; i < 1000; i++ {
		diff := leftList[i] - rightList[i]
		if diff < 0 {
			diff = diff*-1
		}
		distance += diff;
	}

	fmt.Println(distance)
}

func solve_challange_2(leftList []int, rightList []int) {
	// hashmap the leftlist
	rightMap := make(map[int]int) // value : freq
	
	for _, val := range rightList {
		rightMap[val]++
	}

	// iterate over left map and find the freq
	score := 0
	for _, val := range leftList {
		score += (rightMap[val] * val)
	}

	fmt.Println(score)
}

func main() {
	data, err := readFile()
	if err != nil {
		fmt.Println("error reading file", err)
		os.Exit(1)
	}
	leftList, rightList := parseInput(data)
	solve_challange_2(leftList, rightList)
}
