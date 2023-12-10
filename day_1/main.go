package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var Numerics = map[string]string{
	"one":   "one1one",
	"two":   "two2two",
	"three": "three3three",
	"four":  "four4four",
	"five":  "five5five",
	"six":   "six6six",
	"seven": "seven7seven",
	"eight": "eight8eight",
	"nine":  "nine9nine",
}

func matchStringToNumber(input string) string {
	numeric := strings.TrimSpace(input)

	keys := make([]string, 0, len(Numerics))

	for k := range Numerics {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {

		numeric = strings.Replace(numeric, k, Numerics[k], len(input))

	}

	return numeric
}

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

func findFirstNumeric(strings []string) string {
	var numberLike string
out:
	for i := 0; i < len(strings); i++ {
		character := strings[i]
		if isNumeric(character) {
			numberLike = character
			break out
		}
	}
	return numberLike
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Failed to read file", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	var total int64

	for i := 0; i < len(lines); i++ {

		firstNumber, lastNumber := findFirstNumeric(strings.Split(matchStringToNumber(lines[i]), "")), findFirstNumeric(reverseArray(strings.Split(matchStringToNumber(lines[i]), "")))

		fmt.Println(lines[i], firstNumber, lastNumber)
		sum, err := strconv.ParseInt(firstNumber+lastNumber, 10, 64)

		if err != nil {
			fmt.Println("Failed to parse number", err)
		}

		total += sum
	}
	fmt.Println(total)
}
