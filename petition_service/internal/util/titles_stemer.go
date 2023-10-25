package util

import (
	"fmt"
	"strings"

	"github.com/catness812/e-petitions-project/petition_service/internal/util/romanian"
	"github.com/kljensen/snowball"
	"github.com/pemistahl/lingua-go"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func PreprocessText(text string) []string {
	text = strings.NewReplacer(",", "", ".", "", "!", "", "?", "").Replace(text)
	lang, exists := detectLanguage(text)
	language := ""
	words := strings.Fields(text)
	if !exists {
		return words
	}
	if lang == lingua.Romanian {
		language = "romanian"
	} else {
		language = strings.ToLower(fmt.Sprint(lang))
	}
	stemmedWord := ""
	var err error
	var stemmedWords []string
	for _, word := range words {
		if language != "romanian" {
			stemmedWord, err = snowball.Stem(word, language, true)
			if err != nil {
				fmt.Println("Error stemming word:", err)
				stemmedWord = ""
			}
		} else {
			stemmedWord = romanian.Stem(word, false)
			stemmedWords = append(stemmedWords, word)
			// stemmedWord = word
		}

		stemmedWords = append(stemmedWords, stemmedWord)
	}
	// processedText := strings.Join(stemmedWords, " ")
	return stemmedWords
}

func detectLanguage(text string) (lingua.Language, bool) {
	languages := []lingua.Language{lingua.English, lingua.Romanian, lingua.Russian}

	detector := lingua.NewLanguageDetectorBuilder().FromLanguages(languages...).Build()

	language, exists := detector.DetectLanguageOf(text)
	return language, exists
}

func calculateWordSimilarities(words1, words2 []string) float64 {
	similarities := make([]float64, len(words1))
	mx := len(words1)
	if len(words2) > mx {
		mx = len(words2)
	}
	//needs optimization
	for i, word1 := range words1 {
		bestScore := 0.0
		bestWordIndex := -1

		for j, word2 := range words2 {
			if j == bestWordIndex {
				continue
			}
			similarity := calculateSimilarity(word1, word2)
			if similarity > bestScore {
				bestScore = similarity
				bestWordIndex = j
			}
		}

		similarities[i] = bestScore
	}

	return calculateMean(similarities, mx)
}

func calculateSimilarity(text1, text2 string) float64 {
	return levenshtein.RatioForStrings([]rune(text1), []rune(text2), levenshtein.DefaultOptions)
}

func calculateMean(values []float64, mx int) float64 {
	n := len(values)
	if n == 0 {
		return 0.0
	}
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum / float64(mx)
}

func CalculateTitleSimilarity(title1 []string, title2 string) float64 {
	//needs optimization

	// Preprocess the titles
	// processedTitle1 := PreprocessText(title1)
	processedTitle2 := PreprocessText(title2)

	// Calculate word-level similarity
	wordMedianSimilarities := calculateWordSimilarities(title1, processedTitle2)
	return wordMedianSimilarities
}
