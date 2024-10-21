package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func stringToInt(str string) int {
	value, err := strconv.Atoi(str)
	check(err)
	return value
}

func readInput() [][][]int {
	inputFile, err := os.Open("./input.txt")
	check(err)
	defer inputFile.Close()

	result := [][][]int{}
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			continue
		}
		ranges := strings.Split(text, ",")
		range1 := strings.Split(ranges[0], "-")
		range2 := strings.Split(ranges[1], "-")

		range1Int := []int{stringToInt(range1[0]), stringToInt(range1[1])}
		range2Int := []int{stringToInt(range2[0]), stringToInt(range2[1])}

		result = append(result, [][]int{range1Int, range2Int})
	}
	return result
}

func eval1(ranges [][]int) int {
	if ranges[0][0] <= ranges[1][0] && ranges[0][1] >= ranges[1][1] {
		return 1
	}
	if ranges[1][0] <= ranges[0][0] && ranges[1][1] >= ranges[0][1] {
		return 1
	}
	return 0
}

func eval2(ranges [][]int) int {
	if ranges[0][0] <= ranges[1][0] && ranges[0][1] >= ranges[1][1] {
		return 1
	}
	if ranges[1][0] <= ranges[0][0] && ranges[1][1] >= ranges[0][1] {
		return 1
	}
	if ranges[0][0] < ranges[1][0] && ranges[0][1] >= ranges[1][0] {
		return 1
	}
	if ranges[1][0] < ranges[0][0] && ranges[1][1] >= ranges[0][0] {
		return 1
	}
	return 0
}

func main() {
	input := readInput()

	// Part 1:
	sum1 := 0
	for _, line := range input {
		sum1 += eval1(line)
	}
	fmt.Println("Part 1: ", sum1)

	// Part 2:
	sum2 := 0
	for _, line := range input {
		sum2 += eval2(line)
	}
	fmt.Println("Part 2: ", sum2)
}
