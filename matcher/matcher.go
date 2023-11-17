package matcher

import "strings"

// Matcher is the interface that wraps the Match method.
type Matcher interface {
	Match(string, MatchOptions) (bool, error)
}

// Word is a string that represents a word.
type Word string

// Phrase is a list of words that represents a phrase.
type Phrase []Word

// MatchOptions is the options used to match a phrase against a content.
type MatchOptions struct {
	// Phrases is a list of phrases to compare against the content.
	Phrases []Phrase

	// InteruptThreshold is the maximum number of words that can be skipped
	// before the match is interupted.
	InteruptThreshold int

	// SimilarityThreshold is the minimum similarities score required for a match.
	SimilarityThreshold float64

	// Similarity is the function used to compare two words.
	Similarity func(Word, Word) float64
}

// NewMatcher returns a new Matcher.
func NewMatcher() Matcher {
	return &matcher{}
}

// matcher is the default implementation of the Matcher interface.
type matcher struct{}

// Words converts a string into a list of words.
func Words(content string) (words []Word) {
	for _, word := range strings.Fields(content) {
		words = append(words, Word(word))
	}

	return
}

// Match returns true if the content matches the phrases.
func (m *matcher) Match(content string, options MatchOptions) (bool, error) {
	words := Words(content)
	for _, phrase := range options.Phrases {
		if match, err := matchPhrase(words, phrase, options); err != nil {
			return false, err
		} else if match {
			return true, nil
		}
	}

	return false, nil
}

// matchPhrase returns true if the words match the phrase.
func matchPhrase(words []Word, phrase Phrase, options MatchOptions) (bool, error) {
	i, latestInteruptIndex := 0, -1
	for _, word := range words {
		// If the word is similar to the current word in the phrase, we continue.
		if options.Similarity(word, phrase[i]) >= options.SimilarityThreshold {
			i++
		}

		// If the word is not similar to the current word in the phrase, we check
		// if we can skip it.
		if latestInteruptIndex > -1 && latestInteruptIndex+options.InteruptThreshold < i {
			return false, nil
		}

		// If we reached the end of the phrase, we return true.
		if i == len(phrase) {
			return true, nil
		}
	}

	return false, nil
}
