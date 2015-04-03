package str

import (
	"unicode/utf8"

	"github.com/adrg/go-utils/math"
)

func JWDistance(first, second string) float64 {
	fLen := utf8.RuneCountInString(first)
	sLen := utf8.RuneCountInString(second)
	if fLen > sLen {
		first, second = second, first
		fLen, sLen = sLen, fLen
	}

	halfLen := fLen/2 + 1
	fm := MatchingRunesWithLimit(first, second, halfLen)
	sm := MatchingRunesWithLimit(second, first, halfLen)

	fmLen := len(fm)
	smLen := len(sm)
	if fmLen == 0 || smLen == 0 || fmLen != smLen {
		return 0.0
	}

	trans := Transpositions(string(fm), string(sm))
	dist := (float64(fmLen)/float64(fLen) + float64(smLen)/(float64(sLen)) +
		float64(fmLen-trans)/(float64(fmLen))) / 3.0

	cpLen := utf8.RuneCountInString(CommonPrefix(first, second))
	if cpLen > 4 {
		cpLen = 4
	}

	return math.Round((dist+(0.1*float64(cpLen)*(1.0-dist)))*100.0, 2) / 100.0
}
