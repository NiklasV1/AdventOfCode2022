package main

import (
	"bufio"
	"os"
	"slices"
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

type Input struct {
	stacks       []string
	instructions [][]int
}

func readInput() Input {
	inputFile, err := os.Open("./input.txt")
	check(err)
	defer inputFile.Close()

	stacks := []string{}
	instructions := [][]int{}
	scanner := bufio.NewScanner(inputFile)

	instructionsReached := false
	for scanner.Scan() {
		text := scanner.Text()
		if !instructionsReached {
			if text == "" || strings.Contains(text, "1") {
				instructionsReached = true
				continue
			}

			// Parse Map
			text += " "
			runes := []rune(text)
			chunks := slices.Chunk(runes, 4)
			for value := range chunks {
				if value[1] != " " {

				}
			}

		} else {
			// Parse Instructions
			if text == "" {
				continue
			}
			splitText := strings.Split(text, " ")
			instructions = append(instructions, []int{stringToInt(splitText[1]), stringToInt(splitText[3]), stringToInt(splitText[5])})
		}

	}
	result := Input{stacks: stacks, instructions: instructions}
	return result
}
