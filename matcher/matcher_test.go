package matcher

import (
	"fmt"
)

func ExampleMatcher_Match_exactMatch() {
	m := Matcher{
		Phrases:             []Phrase{{"hello", "world"}},
		InteruptThreshold:   1,
		SimilarityThreshold: 1.0,
		Similarity: func(a, b Word) float64 {
			if a == b {
				return 1.0
			}
			return 0.0
		},
	}

	content := "hello world"
	result := m.Match(content)
	fmt.Println(result)
	// Output: true
}

func ExampleMatcher_Match_partialMatchWithInterruption() {
	m := Matcher{
		Phrases:             []Phrase{{"quick", "fox", "over", "lazy", "dog"}},
		InteruptThreshold:   1,
		SimilarityThreshold: 0.8,
		Similarity: func(a, b Word) float64 {
			if a == b {
				return 1.0
			}
			return 0.0
		},
	}

	content := "the quick red fox jumps over the lazy brown dog"
	result := m.Match(content)
	fmt.Println(result)
	// Output: true
}

func ExampleMatcher_Match_noMatch() {
	m := Matcher{
		Phrases:             []Phrase{{"apple", "orange"}},
		InteruptThreshold:   1,
		SimilarityThreshold: 0.8,
		Similarity: func(a, b Word) float64 {
			if a == b {
				return 1.0
			}
			return 0.0
		},
	}

	content := "grapefruit"
	result := m.Match(content)
	fmt.Println(result)
	// Output: false
}
