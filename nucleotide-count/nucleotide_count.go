package dna

import (
	"errors"
	"strings"
)

type Histogram map[rune]int
type DNA string

const nucleotides string = "ACGT"

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{}
	h = setDefaultHistogram(h, nucleotides)
	errorCatch := 0
	for _, v := range nucleotides {
		h[v] = strings.Count(string(d), string(v))
		errorCatch += h[v]
	}
	if errorCatch != len(d) {
		return nil, errors.New("Wrong nucleotide")
	}

	return h, nil
}

func setDefaultHistogram(h Histogram, nucleotides string) Histogram {
	for _, v := range nucleotides {
		h[v] = 0
	}
	return h
}
