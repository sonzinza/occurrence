package occurrence

import (
	"sort"
	"strings"
	"unicode"
)

type WordOccurence struct {
	Word       string `json:"word"`
	Occurrence int    `json:"occurrence"`
}

func GetOccurence(input string) []WordOccurence {
	var tmp []string
	input = strings.ToLower(input)
	wordList := strings.FieldsFunc(input, func(r rune) bool {
		return unicode.IsSpace(r) || r == '.' || r == ',' // this rune allow count occurrence for word include ',' and '.'
	})
	counts := make(map[string]int)
	for _, word := range wordList {
		counts[word] += 1
	}
	res := make([]string, 0, len(counts))

	for word := range counts {
		res = append(res, word)
	}

	sort.Slice(res, func(i, j int) bool {
		return counts[res[i]] > counts[res[j]]
	})

	countResponse := len(res)
	wordOccurrence := make([]WordOccurence, 0, countResponse)
	if countResponse < 10 {
		tmp = res[:countResponse]
	} else {
		tmp = res[:10]
	}
	for _, val := range tmp {
		wordOccurrence = append(wordOccurrence, WordOccurence{
			Word:       val,
			Occurrence: counts[val],
		})
	}
	return wordOccurrence
}
