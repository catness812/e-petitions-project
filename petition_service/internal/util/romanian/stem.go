package romanian

import (
	"strings"

	"github.com/kljensen/snowball/snowballword"
)

var Step1Success bool = false
var Step2Success bool = false

func Stem(word string, stemStopWords bool) string {
	word = strings.ToLower(strings.TrimSpace(word))
	w := snowballword.New(word)

	if len(w.RS) <= 2 || (stemStopWords == false && isStopWord(word)) {
		return word
	}
	preprocess(w)
	// fmt.Println("preprocess:", string(w.RS))
	step0(w)
	// fmt.Println("step0:", string(w.RS))
	step1(w)
	// fmt.Println("step1:", string(w.RS))
	step2(w)
	// fmt.Println("step2:", string(w.RS))
	step3(w)
	// fmt.Println("step3:", string(w.RS))
	step4(w)
	// fmt.Println("step4:", string(w.RS))
	return w.String()
}
