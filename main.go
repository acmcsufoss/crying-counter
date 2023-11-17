// Package main demonstrates a bare simple bot without a state cache. It logs
// all messages it sees into stderr.
package main

import (
	"fmt"

	"github.com/acmcsufoss/crying-counter/matcher"
)

func main() {
	m := matcher.NewMatcher()
	result, err := m.Match("hello world", matcher.MatchOptions{
		Phrases: []matcher.Phrase{
			{
				"hello",
				"world",
			},
		},
		InteruptThreshold:   1,
		SimilarityThreshold: 0.5,
		Similarity: func(a matcher.Word, b matcher.Word) float64 {
			return 1
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
