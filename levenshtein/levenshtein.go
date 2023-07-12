package levenshtein

import (
	"fmt"
	"strings"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

// CheckOptions is the options for the check function.
type CheckOptions struct {
	// Targets is a list of strings to compare against the content.
	Targets [][]string

	// Content is the content or text to search for similarity.
	Content string

	// WindowSize is the size of the sliding window.
	WindowSize int

	// SimilarityThreshold is the minimum similarities score required for a match.
	SimilarityThreshold float64

	// PartialWindowEnabled is a flag indicating whether partial windows are enabled.
	PartialWindowEnabled bool
}

// Check if Content is similar to any of the strings in TargetWindows.
func Check(options CheckOptions) (bool, error) {
	// Split the content into words.
	words := strings.Fields(options.Content)

	// Keep track of the current window in each target window.
	frequencyMap := map[int]int{}
	for i := 0; i < len(options.Targets); i++ {
		frequencyMap[i] = 0
	}

	// Iterate through the words.
	for i := 0; i < len(words); i++ {
		// Iterate through the WindowSize.
		for j := 0; j < options.WindowSize; j++ {
			// Get the current window.
			window := strings.Join(words[i:i+j+1], " ")

			// Check if window is similar to any candidate targets.
			for k, target := range options.Targets {
				targetIndex, ok := frequencyMap[k]
				if !ok {
					// If the target window is not in the frequency map, throw an error.
					return false, fmt.Errorf("target window not in frequency map")
				}
				candidate := target[targetIndex]
				// func RatioForStrings(source []rune, target []rune, op Options) float64
				similarity := levenshtein.RatioForStrings([]rune(window), []rune(candidate), levenshtein.DefaultOptions)

				if similarity >= options.SimilarityThreshold {
					// If the similarity is greater than the threshold, increment the frequency map.
					frequencyMap[k]++
					// If the frequency map is equal to the length of the target window, return true.
					if frequencyMap[k] == len(target) {
						return true, nil
					}
					break
				}

				// Examples:
				// [[“i”, “am”, “crying”], [“im”, “crying”]]

				// ["omg", "i", "am", "literally", "crying"]

				// ["omg"] => ["omg"]
				// ["omg", "i"] => ["omg i"]
				// ["omg", "i", "am"] => ["omg i am"]

				// ["i"] => ["i"]
				// ["i", "am"] => ["i am"]
				// ["i", "am", "literally"] => ["i am literally"]

				// ["am"] => ["am"]
				// ["am", "literally"] => ["am literally"]
				// ["am", "literally", "crying"] => ["am literally crying"]

				// ["literally"] => ["literally"]
				// ["literally", "crying"] => ["literally crying"]

				// ["crying"] => ["crying"]

				// {
				// 	{
				// 	 0: 0
				// 	 1: 0
				// 	 2: 0
				// 	}
				// }

			}
		}
	}
	return false, nil
}
