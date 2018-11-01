package scrabble

import (
	"strings"
)

var scores = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

//Score calculates the value of a scrabble word
func Score(word string) int {
	var total int
	for _, v := range strings.ToUpper(word) {
		for key, value := range scores {
			if strings.Contains(key, string(v)) {
				total += value

			}
		}
	}
	return total
}
