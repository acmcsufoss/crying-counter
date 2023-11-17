// Package main demonstrates a bare simple bot without a state cache. It logs
// all messages it sees into stderr.
package main

import (
	"fmt"

	"github.com/acmcsufoss/crying-counter/matcher"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func main() {
	m := matcher.Matcher{
		Phrases: []matcher.Phrase{
			{
				"hello",
				"world",
			},
		},
		InteruptThreshold:   1,
		SimilarityThreshold: 0.5,
		Similarity: func(a matcher.Word, b matcher.Word) float64 {
			return levenshtein.RatioForStrings([]rune(a), []rune(b), levenshtein.DefaultOptions)
		},
	}
	result := m.Match("hello world")
	fmt.Println(result)
}
