package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var regex = regexp.MustCompile(`[a-zA-Zа-яА-Я]+-?[a-zA-Zа-яА-Я]*`)

func Top(top int, text string) (res []string) {
	matches := regex.FindAllString(text, -1)

	if matches == nil {
		return nil
	}

	wordFrequencies := map[string]int{}
	for _, word := range matches {
		wordFrequencies[strings.ToLower(word)]++
	}

	frequencyWords := map[int][]string{}
	for word, frequency := range wordFrequencies {
		frequencyWords[frequency] = append(frequencyWords[frequency], word)
	}

	frequencies := []int{}
	for frequency := range frequencyWords {
		frequencies = append(frequencies, frequency)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(frequencies)))
	for _, frequency := range frequencies {
		sort.Strings(frequencyWords[frequency])
		res = append(res, frequencyWords[frequency]...)
	}

	if len(wordFrequencies) < top {
		top = len(wordFrequencies)
	}
	return res[:top]
}
