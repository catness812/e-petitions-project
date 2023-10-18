package romanian

import (
	"strings"

	"github.com/kljensen/snowball/snowballword"
)

// STEP 2: Removal of standard suffixes
func step2(word *snowballword.SnowballWord) bool {
	step2Suffixes := []string{
		"abila",
		"abile",
		"abili",
		"abilă",
		"ibila",
		"ibile",
		"ibili",
		"ibilă",
		"atori",
		"itate",
		"itati",
		"ităţi",
		"abil",
		"ibil",
		"oasa",
		"oasă",
		"oase",
		"anta",
		"ante",
		"anti",
		"antă",
		"ator",
		"ităi",
		"iune",
		"iuni",
		"isme",
		"ista",
		"iste",
		"isti",
		"istă",
		"işti",
		"ata",
		"ată",
		"ati",
		"ate",
		"uta",
		"ută",
		"uti",
		"ute",
		"ita",
		"ită",
		"iti",
		"ite",
		"ica",
		"ice",
		"ici",
		"ică",
		"osi",
		"oşi",
		"ant",
		"iva",
		"ive",
		"ivi",
		"ivă",
		"ism",
		"ist",
		"at",
		"ut",
		"it",
		"ic",
		"os",
		"iv",
	}
	Step2Success = false
	for _, suffix := range step2Suffixes {
		if word.HasSuffixRunes([]rune(suffix)) && strings.Contains(word.R2String(), suffix) {
			Step2Success = true
			if suffix == "iune" || suffix == "iuni" {
				if len(word.RS) >= 5 && word.RS[len(word.RS)-5] == 0x021B {
					word.RS[len(word.RS)-5] = 't'
				}

			} else if suffix == "ism" || suffix == "isme" || suffix == "ist" || suffix == "ista" ||
				suffix == "iste" || suffix == "isti" || suffix == "istă" || suffix == "iști" {
				word.ReplaceSuffix(suffix, "ist", false)
			} else {
				word.RemoveFirstSuffix(suffix)
			}
			break
		}
	}

	return Step2Success
}
