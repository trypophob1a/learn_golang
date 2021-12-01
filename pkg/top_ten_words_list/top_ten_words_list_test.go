package top_ten_words_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testSliceResult = []string{"lorem", "ipsum", "dolor", "sit", "amet", "lorem"}
var testMapResult = map[string]int{"amet": 1, "dolor": 1, "ipsum": 1, "lorem": 2, "sit": 1}

func TestNormaliseText(t *testing.T) {

	textInFile := `Lorem! Ipsum! 
			 dolor sit amet Lorem.
	`
	assert.Equal(t, testSliceResult, normaliseText(textInFile))

	simpleText := "Lorem! Ipsum! dolor sit amet Lorem."
	assert.Equal(t, testSliceResult, normaliseText(simpleText))
}

func TestCountingWords(t *testing.T) {

	assert.Equal(t, testMapResult, countingWords(testSliceResult))
}

func TestSortedTop(t *testing.T) {

	var top WordCounterList = []WordCounter{
		{"lorem", 2}, {"amet", 1},
		{"dolor", 1}, {"ipsum", 1},
		{"sit", 1},
	}
	assert.Equal(t, top, sortedTop(testMapResult))
}

func TestTopFirst(t *testing.T) {

	assert.Equal(t, "lorem", sortedTop(testMapResult)[0].Word)
	assert.Equal(t, 2, sortedTop(testMapResult)[0].Count)
}
