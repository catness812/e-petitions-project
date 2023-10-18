package romanian

import (
	"strings"

	"github.com/kljensen/snowball/snowballword"
)

// STEP 1: Reduction of combining suffixes
func step1(word *snowballword.SnowballWord) bool {
	var step1Suffixes = []string{
		"abilitate",
		"abilitati",
		"abilităţi",
		"ibilitate",
		"abilităi",
		"ivitate",
		"ivitati",
		"ivităţi",
		"icitate",
		"icitati",
		"icităţi",
		"icatori",
		"ivităi",
		"icităi",
		"icator",
		"aţiune",
		"atoare",
		"ătoare",
		"iţiune",
		"itoare",
		"iciva",
		"icive",
		"icivi",
		"icivă",
		"icala",
		"icale",
		"icali",
		"icală",
		"ativa",
		"ative",
		"ativi",
		"ativă",
		"atori",
		"ători",
		"itiva",
		"itive",
		"itivi",
		"itivă",
		"itori",
		"iciv",
		"ical",
		"ativ",
		"ator",
		"ător",
		"itiv",
		"itor",
	}

	for {
		replacementDone := false

		for _, suffix := range step1Suffixes {
			if word.HasSuffixRunes([]rune(suffix)) && strings.Contains(word.R1String(), suffix) {
				Step1Success = true
				replacementDone = true
				switch suffix {
				case "abilitate", "abilitati", "abilităi", "abilităţi":
					word.ReplaceSuffix(suffix, "abil", false)
				case "ibilitate":
					word.RemoveFirstSuffix("itate")
				case "ivitate", "ivitati", "ivităi", "ivităţi":
					word.ReplaceSuffix(suffix, "iv", false)
				case "icitate", "icitati", "icităi", "icităţi", "icator", "icatori", "iciv", "iciva", "icive", "icivi", "icivă", "ical", "icala", "icale", "icali", "icală":
					word.ReplaceSuffix(suffix, "ic", false)
				case "aţiune":
					word.RS = ReplaceSuffix(string(word.RS), suffix, "aţ")
				case "ativ", "ativa", "ative", "ativi", "ativă", "atoare", "ator", "atori", "ătoare", "ător", "ători":
					word.ReplaceSuffix(suffix, "at", false)
				case "itiv", "itiva", "itive", "itivi", "itivă", "iţiune", "itoare", "itor", "itori":
					word.ReplaceSuffix(suffix, "it", false)
				}
			}
		}

		if !replacementDone {
			break
		}
	}
	return true
}
