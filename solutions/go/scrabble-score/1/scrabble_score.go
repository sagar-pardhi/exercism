package scrabble

import (
	"strings"
	"unicode/utf8"
)

func Score(word string) int {
	score := 0

	for i := range word {
		runeValue, _ := utf8.DecodeRuneInString(word[i:])

		char := strings.ToUpper(string(runeValue))

		switch char {
		case "A",
			"E",
			"I",
			"O",
			"U",
			"L",
			"N",
			"R",
			"S",
			"T":
			score = score + 1
		case "D", "G":
			score = score + 2
		case "B",
			"C",
			"M",
			"P":
			score = score + 3
		case "F",
			"H",
			"V",
			"W",
			"Y":
			score = score + 4
		case "K":
			score = score + 5
		case "J", "X":
			score = score + 8
		case "Q", "Z":
			score = score + 10
		default:
			score = score + 0
		}
	}

	return score
}
