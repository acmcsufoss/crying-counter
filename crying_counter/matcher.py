"""
SAMPLE USAGE
from matcher import is_similar, Matcher

# Use 1
boolean = is_similar("im crying", ["im crying"], ["i am crying"])
print(boolean)

# Use 2
is_match = Matcher(["im crying", "i am crying"])
assert(is_match("im crying"))
assert(is_match("i am crying"))
"""

import string
from typing import Callable

SIMILARITY_THRESHOLD = 0.8
INTERRUPT_THRESHOLD = 1

# single word with no spaces
Word = str
# multiple words that can have spaces
Phrase = str
# Function that takes two words and returns a float between 0.0 and 1.0
# indicating the similarity score as a percent
SimilarityFunc = Callable[[Word, Word], float]

def default_similarity(a: Word, b: Word) -> float:
    """
    Returns a similarity score based on the levenshtein distance algorithm by default
    """
    def levenshtein_distance(a: Word, b: Word) -> int:
        """
        Levenshtein distance for calculating distance between two words
        Reference: https://en.wikipedia.org/wiki/Levenshtein_distance
        """
        m, n = len(a), len(b)
        dp = [[0] * (n + 1) for _ in range(m + 1)]

        for i in range(m + 1):
            for j in range(n + 1):
                if i == 0:
                    dp[i][j] = j
                elif j == 0:
                    dp[i][j] = i
                elif a[i - 1] == b[j - 1]:
                    dp[i][j] = dp[i - 1][j - 1]
                else:
                    dp[i][j] = 1 + min(dp[i - 1][j], dp[i][j - 1], dp[i - 1][j - 1])

        return dp[m][n]
    distance = levenshtein_distance(a, b)
    # similarity calculation
    similarity = 1.0 - (distance / max(len(a), len(b)))
    return similarity

def match_phrase(
        words: list[Word],
        phrase: Phrase,
        similarity_threshold: float,
        interrupt_threshold: int,
        similarity_func: SimilarityFunc
    ) -> bool:
    """
    Matches a list of words to a single phrase.
    """
    # first split a phrase into words so
    # we can compare word by word
    phrase_words = phrase.split()
    # variables tracking the current word and phrase word in their respective lists
    word_index = 0
    phrase_word_index = 0
    latest_interrupt_index = -1

    while word_index < len(words):
        # one word is similar to another in the phrase
        if similarity_func(words[word_index], phrase_words[phrase_word_index]) >= similarity_threshold:
            # reset interrupt index
            latest_interrupt_index = word_index
            # look for the next word in the phrase
            phrase_word_index += 1
        # too many interrupts, word and phrase up to this point do not match, restart search at the beginning of the phrase
        if latest_interrupt_index > -1 and latest_interrupt_index + interrupt_threshold < word_index:
            word_index = latest_interrupt_index
            latest_interrupt_index = -1
            phrase_word_index = 0
        # all words in the phrase were found which means they match!
        if phrase_word_index == len(phrase_words):
            return True

        word_index += 1

    return False

def to_words(phrase: Phrase) -> list[Word]:
    """
    Splits a given phrase and returns a cleaned list of words.

    Converts every word to lowercase and removes punctuation.
    """
    words = phrase.split()
    cleaned_words = []
    for word in words:
        word = ''.join(c.lower() for c in word if c not in string.punctuation)
        cleaned_words.append(word)
    return cleaned_words

def is_similar(
        phrase_have: Phrase,
        phrases_want: list[Phrase],
        similarity_threshold: float | None = SIMILARITY_THRESHOLD,
        interrupt_threshold: int | None = INTERRUPT_THRESHOLD,
        similarity_func: SimilarityFunc | None = default_similarity
    ) -> bool:
    """
    Matches a phrase with any phrase in a list of phrases.

    Parameters:
        `phrase_have` (`Phrase`): A phrase to match.
        `phrase_want` (`list[Phrase]`): A list of phrases to be matched with.
        `similarity_threshold` (`float`): A float from 0.0 to 1.0 indicating
            the required similarity percent a list of words should match a phrase.
            Default value is 0.8.
        `interrupt_threshold` (`int`): The maximum number of consecutive words that
            could interrupt a phrase while still maintaining a possible match.
            Default value is 1.
        `similarity_func` (`SimilarityFunc`): A similarity function that takes in two words
            and returns a float representing the similar score between them.
            Default implementation is based on the Levenshtein algorithm.
    Returns:
        `bool`: whether the phrase is similar to any phrase in the list.
    """
    words = to_words(phrase_have)
    return any(
        match_phrase(words, phrase, similarity_threshold, interrupt_threshold, similarity_func)
        for phrase in phrases_want
    )

def Matcher(
        phrase_want: list[Phrase],
        similarity_threshold: float | None = SIMILARITY_THRESHOLD,
        interrupt_threshold: int | None = INTERRUPT_THRESHOLD,
        similarity_func: SimilarityFunc | None = default_similarity
    ):
    """
    Creates a Matcher functor which can be called with a given phrase.
    """
    def match(phrase_have: Phrase) -> bool:
        return is_similar(phrase_have, phrase_want, similarity_threshold, interrupt_threshold, similarity_func)
    return match
