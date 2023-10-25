package romanian

import "github.com/kljensen/snowball/snowballword"

func preprocess(word *snowballword.SnowballWord) {
	for i := 1; i < len(word.RS)-1; i++ {
		if isLowerVowel(word.RS[i-1]) && isLowerVowel(word.RS[i+1]) {
			if word.RS[i] == 'u' {
				word.RS[i] = 'U'
			} else if word.RS[i] == 'i' {
				word.RS[i] = 'I'
			}
		}
	}

	r1start, r2start, rvstart := findRegions(word)
	word.R1start = r1start
	word.R2start = r2start
	word.RVstart = rvstart

}
