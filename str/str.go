package str

import (
	"unicode/utf8"

	"github.com/adrg/go-utils/math"
)

func CommonPrefix(first, second string) string {
	if utf8.RuneCountInString(first) > utf8.RuneCountInString(second) {
		first, second = second, first
	}

	var commonLen int
	sRunes := []rune(second)
	for i, r := range first {
		if r != sRunes[i] {
			break
		}

		commonLen++
	}

	return string(sRunes[0:commonLen])
}

func Transpositions(first, second string) int {
	var trans int

	if utf8.RuneCountInString(first) > utf8.RuneCountInString(second) {
		first, second = second, first
	}

	sRunes := []rune(second)
	for index, r := range first {
		if r != sRunes[index] {
			trans++
		}
	}

	return trans / 2
}

func MatchingRunesWithLimit(first, second string, lmt int) []rune {
	common := []rune{}

	sRunes := []rune(second)
	for i, r := range first {
		for j := math.Max(0, i-lmt); j < math.Min(i+lmt, len(sRunes)); j++ {
			if r == sRunes[j] {
				common = append(common, sRunes[j])
				sRunes[j] = 0
				break
			}
		}
	}

	return common
}

func Unique(items []string) []string {
	uniq := make([]string, len(items))

	index := 0
	catalog := map[string]struct{}{}
	for _, item := range items {
		if _, ok := catalog[item]; ok {
			continue
		}

		catalog[item] = struct{}{}
		uniq[index] = item
		index++
	}

	return uniq[0:index]
}
