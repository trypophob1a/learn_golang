package top_ten_words_list

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
	"unicode/utf8"
)

func PrintTop(text string) {
	wordsMap := countingWords(normaliseText(text))
	topWords := sortedTop(wordsMap)
	for i, word := range topWords {
		fmt.Printf("#%d %s => %d \n", i+1, word.Word, word.Count)
	}

}

func sortedTop(wordsMap map[string]int) WordCounterList {

	capacity := len(wordsMap)
	if capacity > 10 {
		capacity = 10
	}
	wl := make(WordCounterList, len(wordsMap))
	index := 0
	for w, c := range wordsMap {
		wl[index] = WordCounter{w, c}
		index++
	}
	sort.Sort(sort.Reverse(wl))
	return wl[0:capacity]
}

func countingWords(words []string) map[string]int {
	wordsMap := make(map[string]int, len(words)/2)
	for _, word := range words {
		if utf8.RuneCountInString(word) > 2 {
			wordsMap[word]++
		}
	}
	return wordsMap
}

func normaliseText(text string) []string {

	regExp := regexp.MustCompile("[^А-Яа-яЁёA-Za-z0-9]+")
	newText := strings.ToLower(strings.Trim(text, " "))
	words := regExp.Split(newText, -1)
	if words[len(words)-1] == "" {
		return words[:len(words)-1]
	}

	return words
}

func GetTextInFile(filePath string) string {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return string(file)
}

type WordCounter struct {
	Word  string
	Count int
}

type WordCounterList []WordCounter

func (w WordCounterList) Len() int { return len(w) }
func (w WordCounterList) Less(i, j int) bool {
	if w[i].Count == w[j].Count {
		return w[i].Word > w[j].Word
	}
	return w[i].Count < w[j].Count
}
func (w WordCounterList) Swap(i, j int) { w[i], w[j] = w[j], w[i] }
