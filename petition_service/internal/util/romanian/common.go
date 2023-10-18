package romanian

import (
	"strings"

	"github.com/kljensen/snowball/romance"
	"github.com/kljensen/snowball/snowballword"
)

func isStopWord(word string) bool {
	switch word {
	case "a", "abia", "acea", "aceasta", "această", "aceea", "aceeasi", "acei", "aceia",
		"acel", "acela", "acelasi", "acele", "acelea", "acest", "acesta", "aceste", "acestea",
		"acestei", "acestia", "acestui", "aceşti", "aceştia", "acolo", "acord", "acum", "adica",
		"ai", "aia", "aibă", "aici", "aiurea", "al", "ala", "alaturi", "ale", "alea", "alt", "alta",
		"altceva", "altcineva", "alte", "altfel", "alti", "altii", "altul", "am", "anume", "apoi",
		"ar", "are", "as", "asa", "asemenea", "asta", "astazi", "astea", "astfel", "astăzi", "asupra",
		"atare", "atat", "atata", "atatea", "atatia", "ati", "atit", "atita", "atitea", "atitia", "atunci",
		"au", "avea", "avem", "aveţi", "avut", "azi", "aş", "aşadar", "aţi", "b", "ba", "bine", "bucur",
		"bună", "c", "ca", "cam", "cand", "capat", "care", "careia", "carora", "caruia", "cat", "catre",
		"caut", "ce", "cea", "ceea", "cei", "ceilalti", "cel", "cele", "celor", "ceva", "chiar", "ci", "cinci",
		"cind", "cine", "cineva", "cit", "cita", "cite", "citeva", "citi", "citiva", "conform", "contra", "cu",
		"cui", "cum", "cumva", "curând", "curînd", "când", "cât", "câte", "câţi", "cînd", "cît", "cîte", "câtva",
		"cîţi", "că", "căci", "cărei", "căror", "cărui", "către", "d", "da", "daca", "dacă", "dar", "dat", "datorită", "dată",
		"dau", "de", "deasupra", "deci", "decit", "degraba", "deja", "deoarece", "departe", "desi", "despre", "deşi", "din",
		"dinaintea", "dintr", "dintr-", "dintre", "doar", "doi", "doilea", "două", "drept", "dupa", "după", "dă", "e", "ea",
		"ei", "el", "ele", "era", "eram", "este", "eu", "exact", "eşti", "f", "face", "fara", "fata", "fel", "fi", "fie", "fiecare",
		"fii", "fim", "fiu", "fiţi", "foarte", "fost", "frumos", "fără", "g", "geaba", "graţie", "h", "halbă", "i", "ia", "iar", "ieri",
		"ii", "il", "imi", "in", "inainte", "inaintea", "încotro", "încât", "încît", "între", "întrucât", "întrucît", "îţi", "ăla", "ălea",
		"ăsta", "ăstea", "ăştia", "şapte", "şase", "şi", "ştiu", "ţi", "ţie":
		return true
	}
	return false
}

func isLowerVowel(r rune) bool {
	switch r {
	case 'a', 'ă', 'â', 'e', 'i', 'î', 'o', 'u':
		return true
	}
	return false
}

func ReplaceSuffix(word, oldSuffix, newSuffix string) []rune {
	if strings.HasSuffix(word, oldSuffix) {
		return []rune(word[:len(word)-len(oldSuffix)] + newSuffix)
	}
	return []rune(word)
}

func removeSuffix(word, suffix string) []rune {
	if strings.HasSuffix(word, suffix) {
		return []rune(word[:len(word)-len(suffix)])
	}
	return []rune(word)
}

func findRegions(word *snowballword.SnowballWord) (r1start, r2start, rvstart int) {

	r1start = romance.VnvSuffix(word, isLowerVowel, 0)
	r2start = romance.VnvSuffix(word, isLowerVowel, r1start)
	rvstart = len(word.RS)

	if len(word.RS) >= 3 {
		switch {

		case !isLowerVowel(word.RS[1]):

			// If the second letter is a consonant, RV is the region after the
			// next following vowel.
			for i := 2; i < len(word.RS); i++ {
				if isLowerVowel(word.RS[i]) {
					rvstart = i + 1
					break
				}
			}

		case isLowerVowel(word.RS[0]) && isLowerVowel(word.RS[1]):

			// Or if the first two letters are vowels, RV
			// is the region after the next consonant.
			for i := 2; i < len(word.RS); i++ {
				if !isLowerVowel(word.RS[i]) {
					rvstart = i + 1
					break
				}
			}
		default:

			// Otherwise (consonant-vowel case) RV is the region after the
			// third letter. But RV is the end of the word if these
			// positions cannot be found.
			rvstart = 3
		}
	}
	return
}
