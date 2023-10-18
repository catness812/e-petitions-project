package romanian

import (
	"strings"

	"github.com/kljensen/snowball/snowballword"
)

// STEP 4: Removal of final vowel
func step4(word *snowballword.SnowballWord) bool {
	var replacements = map[string]string{
		"I": "i",
		"U": "u",
	}

	// List of suffixes
	suffixes := []string{"ie", "a", "e", "i", "ă"}
	for _, suffix := range suffixes {
		if word.HasSuffixRunes([]rune(suffix)) {
			if strings.HasSuffix(word.RVString(), suffix) {
				word.RS = removeSuffix(string(word.RS), suffix)
			}
			break
		}
	}
	// Apply replacements
	for old, new := range replacements {
		word.RS = []rune(strings.ReplaceAll(string(word.RS), old, new))
	}
	Step1Success = false
	Step2Success = false
	return true
}
