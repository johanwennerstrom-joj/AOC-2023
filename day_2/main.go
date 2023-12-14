package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var gamePattern = regexp.MustCompile(`Game (\d+):`)

var numberPattern = regexp.MustCompile(`(\d+)`)
var colorPattern = regexp.MustCompile(`\s*([a-zA-Z]+)`)

var MaxValues = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func extractId(dirty string) (string, string) {
	match := gamePattern.FindString(dirty)
	idMatches := numberPattern.FindString(match)
	if len(idMatches) == 0 {
		return "", ""
	}
	id := strings.TrimSpace(idMatches)
	clean := strings.TrimSpace(gamePattern.ReplaceAllString(dirty, ""))
	clean = strings.ReplaceAll(clean, " ", "")

	return id, clean
}

func gamePossible(scoreString string) bool {

	isPossible := true

	scoreString = strings.TrimSpace(scoreString)
	rounds := strings.Split(scoreString, ";")

out:
	for _, round := range rounds {

		roundItems := strings.Split(round, ",")

		for _, item := range roundItems {

			number, color := numberPattern.FindString(item), colorPattern.FindString(item)

			parsedNumber, err := strconv.Atoi(strings.TrimSpace(number))

			if err != nil {
				fmt.Println(err)
			}

			if parsedNumber > MaxValues[color] {
				isPossible = false
				break out
			}
		}

	}

	return isPossible

}

func getPower(scoreString string) int {

	scoreString = strings.TrimSpace(scoreString)
	rounds := strings.Split(scoreString, ";")

	MaxGame := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, round := range rounds {

		roundItems := strings.Split(round, ",")

		for _, item := range roundItems {

			number, color := numberPattern.FindString(item), colorPattern.FindString(item)

			parsedNumber, err := strconv.Atoi(strings.TrimSpace(number))

			if err != nil {
				fmt.Println(err)
			}

			if parsedNumber > MaxGame[color] {
				MaxGame[color] = parsedNumber
			}

		}

	}

	sum := MaxGame["red"] * MaxGame["green"] * MaxGame["blue"]

	return sum

}

func part1and2(games []string) (int, int) {

	var sum, totalPower int

	for _, game := range games {

		id, clean := extractId(game)
		totalPower += getPower(clean)

		if gamePossible(clean) {
			parsedId, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Failed to parse id")
				continue
			}
			sum += parsedId
		}
	}

	return sum, totalPower
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Failed to read file", err)
		return
	}

	games := strings.Split(string(data), "\n")

	part1Solution, part2Solution := part1and2(games)

	fmt.Printf("Part 1 is: %d\nPart 2 is: %d\n", part1Solution, part2Solution)

}
