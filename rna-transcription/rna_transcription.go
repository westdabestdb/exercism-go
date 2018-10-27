package strand

//ToRNA takes a dna type string, and returns rna string.
func ToRNA(dna string) string {
	var dnaLetters = []rune{'A', 'T', 'C', 'G'}
	var rnaLetters = []rune{'U', 'A', 'G', 'C'}
	var returnedDna string
	for _, value := range dna {
		for i, v := range dnaLetters {
			if value == v {
				returnedDna += string(rnaLetters[i])
			}
		}

	}
	return returnedDna
}
