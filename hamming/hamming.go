package hamming

//imports the 'errors' package.
import (
	"errors"
)

func Distance(a, b string) (int, error) {
	var distance int      //defines distance variable as type integer and default value is 0
	if len(a) != len(b) { //checks for the lengths of two dna strands
		return -1, errors.New("strands must be of the same length") //if not same, returns the distance as -1 and returns new error message
	}
	for i := range a { //run loop for the range(from 0 to length) of dna strand 'a'
		if a[i] != b[i] { //compares for the non-same nucleotides
			distance++ //if matches, increases the distance by one.
		}
	}
	return distance, nil //returns the distance but no error.
}
