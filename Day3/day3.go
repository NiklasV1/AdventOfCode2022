package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() []string {
	inputFile, err := os.Open("./input.txt")
	check(err)
	defer inputFile.Close()

	result := []string{}
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			continue
		}
		result = append(result, text)
	}
	return result
}

func getPriority(letter rune) int {
	if strings.ToLower(string(letter)) == string(letter) {
		priority := (letter - 'a') + 1
		return int(priority)
	} else {
		priority := (strings.ToLower(string(letter))[0] - 'a') + 27
		return int(priority)
	}
}

func eval1(line string) int {
	part1 := line[:len(line)/2]
	part2 := line[len(line)/2:]

	for _, char := range part1 {
		if strings.Contains(part2, string(char)) {
			return getPriority(char)
		}
	}
	return 0
}

func eval2(lines []string) int {
	for _, char := range lines[0] {
		if strings.Contains(lines[1], string(char)) && strings.Contains(lines[2], string(char)) {
			return getPriority(char)
		}
	}
	return 0
}

func main() {
	input := readInput()

	// Part 1
	sum1 := 0
	for _, line := range input {
		sum1 += eval1(line)
	}
	fmt.Println("Part 1: ", sum1)

	// Part 2
	sum2 := 0
	for line := range slices.Chunk(input, 3) {
		sum2 += eval2(line)
	}
	fmt.Println("Part 2: ", sum2)
}
