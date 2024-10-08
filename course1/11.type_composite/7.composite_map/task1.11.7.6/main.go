package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Word struct {
	Word string
	Pos  int
}

func main() {
	text := "Внимание! Внимание! Покупай срочно срочно крипту только у нас! Биткоин лайткоин эфир по низким ценам! Беги, беги, успевай стать финансово независимым с помощью крипты! Крипта будущее финансового мира!"
	censorMap := map[string]string{"крипта": "фрукты", "крипту": "фрукты", "крипты": "фруктов", "биткоин": "яблоки", "лайткоин": "яблоки", "эфир": "яблоки"}
	filteredText := filterWords(text, censorMap)
	fmt.Println(filteredText)
}

func filterWords(text string, censorMap map[string]string) string {
	sentences := splitSentences(text)

	if len(sentences) > 1 {
		for i, sentence := range sentences {
			sentences[i] = filterWords(sentence, censorMap)
		}
		return strings.Join(sentences, " ")
	}

	words := strings.Fields(text)

	uniqueWords := make(map[string]Word)

	for i, word := range words {
		lowerWord := strings.ToLower(word)

		if replacement, exists := censorMap[lowerWord]; exists {
			words[i] = CheckUpper(word, replacement)
		}

		if _, exists := uniqueWords[lowerWord]; !exists {
			uniqueWords[lowerWord] = Word{Word: word, Pos: i}
		} else {
			words[uniqueWords[lowerWord].Pos] = ""
			uniqueWords[lowerWord] = Word{Word: word, Pos: i}
		}
	}

	return WordsToSentence(words)
}

func WordsToSentence(words []string) string {
	filtered := make([]string, 0, len(words))

	for _, word := range words {
		if word != "" {
			filtered = append(filtered, word)
		}
	}

	return strings.ReplaceAll(strings.Join(filtered, " ")+"!", "!!", "!")
}

func CheckUpper(old, new string) string {
	if len(old) == 0 || len(new) == 0 {
		return new
	}

	chars := []rune(old)

	if unicode.IsUpper(chars[0]) {
		runes := []rune(new)
		new = string(append([]rune{unicode.ToUpper(runes[0])}, runes[1:]...))
	}

	return new
}

func splitSentences(message string) []string {
	originSentensec := strings.Split(message, "!")
	var orphan string
	var sentences []string

	for i, sentence := range originSentensec {
		words := strings.Split(sentence, " ")

		if len(words) == 1 {
			if len(orphan) > 0 {
				orphan += " "
			}

			orphan += words[0] + "!"
			continue
		}

		if orphan != "" {
			originSentensec[i] = strings.Join([]string{orphan, " ", sentence}, " ") + "!"
			orphan = ""
		}

		sentences = append(sentences, originSentensec[i])
	}

	return sentences
}
