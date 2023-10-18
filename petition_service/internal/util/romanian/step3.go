package romanian

import (
	"strings"

	"github.com/kljensen/snowball/snowballword"
)

// STEP 3: Removal of verb suffixes
func step3(word *snowballword.SnowballWord) bool {
	step3Suffixes := []string{
		"seserăţi",
		"aserăți",
		"iserăți",
		"âserăți",
		"userăți",
		"seserăm",
		"aserăm",
		"iserăm",
		"âserăm",
		"userăm",
		"serăți",
		"seseşi",
		"seseră",
		"ească",
		"arăţi",
		"urăţi",
		"irăţi",
		"ârăţi",
		"aseşi",
		"aseră",
		"iseşi",
		"iseră",
		"âseşi",
		"âseră",
		"useşi",
		"useră",
		"serăm",
		"sesem",
		"indu",
		"âindu",
		"ându",
		"ează",
		"eşti",
		"eşte",
		"ăşti",
		"ăşte",
		"eaţi",
		"iaţi",
		"arăm",
		"urăm",
		"irăm",
		"ârăm",
		"asem",
		"isem",
		"âsem",
		"usem",
		"seşi",
		"seră",
		"sese",
		"are",
		"ere",
		"ire",
		"âre",
		"ind",
		"ând",
		"eze",
		"ezi",
		"esc",
		"ăsc",
		"eam",
		"eai",
		"eau",
		"iam",
		"iai",
		"iau",
		"aşi",
		"ară",
		"uşi",
		"ură",
		"işi",
		"iră",
		"âşi",
		"âră",
		"ase",
		"ise",
		"âse",
		"use",
		"aţi",
		"eţi",
		"iţi",
		"âţi",
		"sei",
		"ez",
		"am",
		"ai",
		"au",
		"ea",
		"ia",
		"ui",
		"âi",
		"ăm",
		"em",
		"im",
		"âm",
		"se",
	}
	if !Step1Success && !Step2Success {
		for _, suffix := range step3Suffixes {
			if word.HasSuffixRunes([]rune(suffix)) {
				if strings.HasSuffix(word.RVString(), suffix) {
					if strings.Contains(
						"seserăţi seserăm serăţi seseşi seseră serăm sesem  em seşi seră sese aţi eţi iţi âţi sei ăm em im âm se ",
						suffix) {
						word.RS = removeSuffix(string(word.RS), suffix)
					} else {
						if containsAny(word.RVString(), "ăindu", "âindu", "ăire", "âire", "âind", "ăind", "serăți", "țiune") {
							break
						}
						suffixIndex := strings.Index(word.RVString(), suffix)
						if suffixIndex >= 1 && !strings.HasPrefix(word.RVString(), suffix) &&
							!(strings.Contains("aeio\u0103\xE2\xEE", string(word.RVString()[suffixIndex-1]))) {
							word.RS = removeSuffix(string(word.RS), suffix)
						}
						break
					}
				}
			}
		}

	}
	return true
}

func containsAny(input string, substrings ...string) bool {
	for _, substring := range substrings {
		if strings.Contains(input, substring) {
			return true
		}
	}
	return false
}
