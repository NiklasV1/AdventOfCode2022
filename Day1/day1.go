package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	inputFile, err := os.Open("./input.txt")
	check(err)
	defer inputFile.Close()

	result := []int{}
	scanner := bufio.NewScanner(inputFile)

	sum := 0
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			result = append(result, sum)
			sum = 0
			continue
		}

		value, err := strconv.Atoi(text)
		check(err)

		sum += value
	}

	if sum != 0 {
		result = append(result, sum)
	}

	slices.Sort(result)

	// Part 1:
	fmt.Println("Part 1: ", result[len(result)-1])

	// Part 2:
	topThree := result[len(result)-1] + result[len(result)-2] + result[len(result)-3]
	fmt.Println("Part 2: ", topThree)
}
