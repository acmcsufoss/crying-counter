package matcher

import "strings"

// Matcher is the interface that wraps the Match method.
type Matcher struct {
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

// Word is a string that represents a word.
type Word string

// Phrase is a list of words that represents a phrase.
type Phrase []Word

// Words converts a string into a list of words.
func Words(content string) (words []Word) {
	for _, word := range strings.Fields(content) {
		words = append(words, Word(word))
	}

	return
}

// Match returns true if the content matches the phrases.
func (m *Matcher) Match(content string) bool {
	words := Words(content)
	for _, phrase := range m.Phrases {
		if m.MatchPhrase(words, phrase) {
			return true
		}
	}

	return false
}

// matchPhrase returns true if the words match the phrase.
func (m *Matcher) MatchPhrase(words []Word, phrase Phrase) bool {
	wordIndex, indexInPhase, latestInteruptIndex := 0, 0, -1
	for wordIndex < len(words) {
		// If the word is similar to the current word in the phrase, we continue.
		if m.Similarity(words[wordIndex], phrase[indexInPhase]) >= m.SimilarityThreshold {
			latestInteruptIndex = wordIndex
			indexInPhase++
		}

		// If the word is not similar to the current word in the phrase, we check
		// if we can skip it.
		if latestInteruptIndex > -1 && latestInteruptIndex+m.InteruptThreshold < wordIndex {
			wordIndex, indexInPhase, latestInteruptIndex = latestInteruptIndex, 0, -1
		}

		// If we reached the end of the phrase, we return true.
		if indexInPhase == len(phrase) {
			return true
		}

		wordIndex++
	}

	return false
}
