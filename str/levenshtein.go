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

func LevDistance(first, second string, opts *LevOptions) int {
	if first == second {
		return 0
	}

	fLen := utf8.RuneCountInString(first)
	sLen := utf8.RuneCountInString(second)
	if math.Min(fLen, sLen) == 0 {
		return math.Max(fLen, sLen)
	}
	if fLen > sLen {
		first, second = second, first
		fLen, sLen = sLen, fLen
	}

	if opts == nil {
		opts = defaultLevOpts
	}
	if !opts.CaseSensitive {
		first = strings.ToLower(first)
		second = strings.ToLower(second)
	}

	prevCol := make([]int, sLen+1)
	for i := 0; i < len(prevCol); i++ {
		prevCol[i] = i
	}

	col := make([]int, sLen+1)
	for i := 0; i < fLen; i++ {
		col[0] = i + 1
		for j := 0; j < sLen; j++ {
			delCost := prevCol[j+1] + opts.DelCost
			insCost := col[j] + opts.InsCost
			subCost := prevCol[j]
			if first[i] != second[j] {
				subCost += opts.SubCost
			}

			col[j+1] = math.Min(delCost, insCost, subCost)
		}

		col, prevCol = prevCol, col
	}

	return prevCol[sLen]
}

func LevRatio(first, second string, opts *LevOptions) float64 {
	if first == "" || second == "" {
		return 100
	}

	maxLen := float64(math.Max(len(first), len(second)))
	return (1 - float64(LevDistance(first, second, opts))/maxLen) * 100
}
