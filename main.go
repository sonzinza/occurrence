package main

import (
	"net/http"
	"sort"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
)

type WordOccurence struct {
	Word       string `json:"word"`
	Occurrence int    `json:"occurrence"`
}

func main() {
	r := setUpRouter()
	r.Run(":8081")
}

func getOccurence(c *gin.Context) {
	var inputRequest map[string]string
	if err := c.ShouldBindJSON(&inputRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := strings.ToLower(inputRequest["text"])
	words := strings.FieldsFunc(input, func(r rune) bool {
		return unicode.IsSpace(r) || r == '.' || r == ',' // this rune allow count occurrence for word include ',' and '.'
	})
	wordOccurrence := getWordFrequency(words)
	c.IndentedJSON(http.StatusOK, wordOccurrence)
}

func setUpRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/occurrence", getOccurence)
	return r
}

func getWordFrequency(words []string) []WordOccurence {
	counts := make(map[string]int)

	for _, word := range words {
		counts[word] += 1
	}
	res := make([]string, 0, len(counts))

	for word := range counts {
		res = append(res, word)
	}
	return sortWordFrequency(res, counts)
}

func sortWordFrequency(words []string, wordCounts map[string]int) []WordOccurence {
	var tmp []string
	sort.Strings(words)
	sort.Slice(words, func(i, j int) bool {
		return wordCounts[words[i]] > wordCounts[words[j]]
	})

	countResponse := len(words)
	wordOccurrence := make([]WordOccurence, 0, countResponse)
	if countResponse < 10 {
		tmp = words[:countResponse]
	} else {
		tmp = words[:10]
	}
	for _, val := range tmp {
		wordOccurrence = append(wordOccurrence, WordOccurence{
			Word:       val,
			Occurrence: wordCounts[val],
		})
	}
	return wordOccurrence
}
