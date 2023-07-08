package levenshtein

import (
	"strings"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

// CheckOptions is the options for the check function.
type CheckOptions struct {
     	 //our temporary variable to hold our window of strings
	 //these temporary variables can be made of content or targen depending on which is longer the longer one will be temp
	 //these strings lengths are window siz whic is the minumum size of target compared to content
     	TargetWindows         [][]string
	//user input
	Content            string
	//size of window which can be anywhere from 1 to min(len(content),len(target)) but should be min of those
	WindowSize           int
	//arbitrary look into what it should be tells if we have a close enough match
	SimilarityThreshold  float64
	//idk what this is
	PartialWindowEnabled bool
}

// Check if Content is similar to any of the strings in TargetWindows.
func Check(options CheckOptions) bool {
     	// split user input from object of type struc by spaces and saving it
     	words := strings.Fields(options.Content)
}