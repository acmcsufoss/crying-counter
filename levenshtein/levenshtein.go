package levenshtein

import (
	"strings"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

// CheckOptions is the options for the check function.
type CheckOptions struct {
	TargetWindows         [][]string
	Content            string
	WindowSize           int
	SimilarityThreshold  float64
	PartialWindowEnabled bool
}

// Check if Content is similar to any of the strings in TargetWindows.
func Check(options CheckOptions) bool {
	words := strings.Fields(options.Content) // Split the message content by spaces
}