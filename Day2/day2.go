package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() [][]string {
	inputFile, err := os.Open("./input.txt")
	check(err)
	defer inputFile.Close()

	result := [][]string{}
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			continue
		}
		result = append(result, strings.Split(text, " "))
	}
	return result
}

func eval1(play []string) int {
	switch play[0] {
	case "A":
		switch play[1] {
		case "X":
			return 1+3
		case "Y":
			return 2+6
		case "Z":
			return 3+0
		}
	case "B":
		switch play[1] {
		case "X":
			return 1+0
		case "Y":
			return 2+3
		case "Z":
			return 3+6
		}
	case "C":
		switch play[1] {
		case "X":
			return 1+6
		case "Y":
			return 2+0
		case "Z":
			return 3+3
		}
	}
	return 0
}

func eval2(play []string) int {
	switch play[0] {
	case "A":
		switch play[1] {
		case "X":
			return 0+3
		case "Y":
			return 3+1
		case "Z":
			return 6+2
		}
	case "B":
		switch play[1] {
		case "X":
			return 0+1
		case "Y":
			return 3+2
		case "Z":
			return 6+3
		}
	case "C":
		switch play[1] {
		case "X":
			return 0+2
		case "Y":
			return 3+3
		case "Z":
			return 6+1
		}
	}
	return 0
}

func main() {
	input := readInput()

	// Part 1
	sum1 := 0
	for _, value := range input {
		sum1 += eval1(value)
	}
	fmt.Println("Part 1: ", sum1)

	// Part 2
	sum2 := 0
	for _, value := range input {
		sum2 += eval2(value)
	}
	fmt.Println("Part 2: ", sum2)
}
