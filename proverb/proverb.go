package proverb

import (
	"fmt"
)

// Proverb creates proverbs from slice or rhymes.
func Proverb(rhyme []string) []string {
	var rhymes []string

	if len(rhyme) <= 0 {
		return rhymes
	}

	for i := 0; i+1 < len(rhyme); i++ {
		line := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		rhymes = append(rhymes, line)
	}
	rhymes = append(rhymes, fmt.Sprintf(
		"And all for the want of a %s.",
		rhyme[0]))
	return rhymes
}
