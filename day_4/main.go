package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func convertToSlice(input string) []int {
	var numbers []int
	numberStrings := strings.Fields(input)

	for _, numStr := range numberStrings {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	return numbers
}

func compareNumberSets(numbers []int, winningNumbers []int) []int {
	var common []int

	for _, num := range numbers {
		for _, winningNum := range winningNumbers {
			if num == winningNum {
				common = append(common, num)
				break
			}
		}

	}

	return common
}

func calculatePoints(matches int) int {
	return int(math.Pow(2, float64(matches-1)))
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(data), "\n")

	var total int

	for _, line := range lines {
		score := strings.Split(line, ":")[1]
		formatted := strings.Split(score, "|")

		numbers, winningNumbers := convertToSlice(formatted[0]), convertToSlice(formatted[1])

		common := compareNumberSets(numbers, winningNumbers)

		sum := calculatePoints(len(common))

		total += sum
	}

	fmt.Println(total)
}
