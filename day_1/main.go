package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isNumeric(str string) bool {
	return regexp.MustCompile(`\d`).MatchString(str)
}

func reverseArray(arr []string) []string {
	toRev := arr
	for i, j := 0, len(toRev)-1; i < j; i, j = i+1, j-1 {
		toRev[i], toRev[j] = toRev[j], toRev[i]
	}

	return toRev
}

func findFirstNumeric(stringMap []string) string {
	var numberLike string
out:
	for i := 0; i < len(stringMap); i++ {
		character := stringMap[i]

		if isNumeric(character) {
			numberLike = character
			break out
		}
	}
	return numberLike
}

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Failed to read file", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	var total int64

	for i := 0; i < len(lines); i++ {

		firstNumber := findFirstNumeric(strings.Split(lines[i], ""))
		secondNumber := findFirstNumeric(reverseArray(strings.Split(lines[i], "")))

		sum, err := strconv.ParseInt(firstNumber+secondNumber, 10, 0)

		if err != nil {
			fmt.Println("Failed to parse number", err)
		}

		total += sum
	}
	fmt.Println(total)
}
