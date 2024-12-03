package main

func isPrefixOfWord(sentence string, searchWord string) int {
	wordOrder := 1
	i := 0
	matchIndex := 0

	for i < len(sentence) {
		if sentence[i] == ' ' {
			wordOrder++
			matchIndex = 0
		} else if sentence[i] == searchWord[matchIndex] {
			matchIndex++
			if matchIndex == len(searchWord) {
				return wordOrder
			}
		} else {
			for i < len(sentence) && sentence[i] != ' ' {
				i++
			}
			wordOrder++
			matchIndex = 0
		}
		i++
	}

	return -1
}
