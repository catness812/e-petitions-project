package romanian

import (
	"strings"

	"github.com/kljensen/snowball/snowballword"
)

// STEP 0: Removal of plurals and other simplifications
func step0(word *snowballword.SnowballWord) bool {
	var step0Suffixes = []string{
		"iilor",
		"ului",
		"elor",
		"iile",
		"ilor",
		"atei",
		"aţie",
		"aţia",
		"aua",
		"ele",
		"iua",
		"iei",
		"ile",
		"ul",
		"ea",
		"ii",
	}

	for _, suffix := range step0Suffixes {
		if word.HasSuffixRunes([]rune(suffix)) && strings.Contains(word.R1String(), suffix) {

			switch suffix {
			case "ul", "ului":
				word.RemoveFirstSuffix(suffix)
			case "aua", "atei", "ile":
				if suffix == "ile" && string(word.RS)[len(word.RS)-5:len(word.RS)-3] != "ab" {
					word.RemoveFirstSuffix(suffix)
				}
			case "ea", "ele", "elor":
				word.ReplaceSuffix(suffix, "e", false)
			case "ii", "iua", "iei", "iile", "iilor", "ilor":
				word.ReplaceSuffix(suffix, "i", false)
			case "aţie", "aţia":
				word.RS = word.RS[:len(word.RS)-1]
			}
			break
		}
	}

	return true
}
