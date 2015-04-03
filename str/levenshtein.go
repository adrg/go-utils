package str

import (
	"strings"
	"unicode/utf8"

	"github.com/adrg/go-utils/math"
)

type LevOptions struct {
	CaseSensitive bool
	InsCost       int
	DelCost       int
	SubCost       int
}

var defaultLevOpts = &LevOptions{true, 1, 1, 1}

func LevDistance(s, t string, opts *LevOptions) int {
	if s == t {
		return 0
	}

	sLen := utf8.RuneCountInString(s)
	tLen := utf8.RuneCountInString(t)
	if math.Min(sLen, tLen) == 0 {
		return math.Max(sLen, tLen)
	}

	if opts == nil {
		opts = defaultLevOpts
	}

	if !opts.CaseSensitive {
		s = strings.ToLower(s)
		t = strings.ToLower(t)
	}

	if sLen > tLen {
		s, t = t, s
		sLen, tLen = tLen, sLen
	}

	prevCol := make([]int, tLen+1)
	for i := 0; i < len(prevCol); i++ {
		prevCol[i] = i
	}

	col := make([]int, tLen+1)
	for i := 0; i < sLen; i++ {
		col[0] = i + 1
		for j := 0; j < tLen; j++ {
			delCost := prevCol[j+1] + opts.DelCost
			insCost := col[j] + opts.InsCost
			subCost := prevCol[j]
			if s[i] != t[j] {
				subCost += opts.SubCost
			}

			col[j+1] = math.Min(delCost, insCost, subCost)
		}

		col, prevCol = prevCol, col
	}

	return prevCol[tLen]
}

func LevRatio(s, t string, opts *LevOptions) float64 {
	maxLen := float64(math.Max(len(s), len(t)))
	return (1 - float64(LevDistance(s, t, opts))/maxLen) * 100
}
