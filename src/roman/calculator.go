package roman

import (
	"fmt"
	"strings"
	"sort"
)

const (
	I = rune('I')
	V = rune('V')
	X = rune('X')
	L = rune('L')
	C = rune('C')
	D = rune('D')
	M = rune('M')
)

var orderedNumeral = map[rune]int{
	I: 0,
	V: 1,
	X: 3,
	L: 4,
	C: 5,
	D: 6,
	M: 7,
}

type Numerals struct {
	value string
}

// Rules
// 1) At-most 1 subtractive prefix
// 2) IV, IX, XL, XC, CD, CM -> 4, 9, 40, 90, 400, 900 only possible subtractive prefixes <= 10% between digits
// 3) At-most 3 repeated additive suffixes
func NewNumerals(numeral string) (Numerals, error) {
	nmls := make([]rune, len(numeral))
	for idx, char := range numeral {
		switch char {
		case I, V, X, L, C, D, M:
			nmls[idx] = char
		default:
			return Numerals{""}, fmt.Errorf("%v is not a valid roman numeral", numeral)
		}
	}
	return Numerals{string(nmls)}, nil
}

func (n1 Numerals) Add(n2 Numerals) Numerals {
	appendedNumerals := Numerals{n1.value + n2.value}
	return normalize(appendedNumerals)
}

func normalize(n Numerals) Numerals {
	subtractiveToAdditiveSuffix := subtractivePrefixToAdditiveSuffixNormalization(n)
	return normalizeAdditiveSuffix(subtractiveToAdditiveSuffix)
}

// Order of normalization matters
func normalizeAdditiveSuffix(n Numerals) Numerals {
	sortedNumerals := sortNumerals(n)
	higherNumeralNormalized := additiveSuffixToSingleSuffixNormalization(sortedNumerals)
	return additiveSuffixToSubtractivePrefixNormalization(higherNumeralNormalized)
}

type sortableNumerals []rune

func (s sortableNumerals) Len() int {
	return len(s)
}

func (s sortableNumerals) Less(i, j int) bool {
	return orderedNumeral[s[j]] <  orderedNumeral[s[i]]
}

func (s sortableNumerals) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortNumerals(n Numerals) Numerals {
	nRunes := sortableNumerals(n.value)
	sort.Sort(nRunes)
	return Numerals{string(nRunes)}
}

func subtractivePrefixToAdditiveSuffixNormalization(n Numerals) Numerals {
	rewriteIVtoIIII := strings.Replace(n.value, "IV", "IIII", -1)
	rewriteIXtoVIIII := strings.Replace(rewriteIVtoIIII, "IX", "VIIII", -1)
	rewriteXLtoXXXX := strings.Replace(rewriteIXtoVIIII, "XL", "XXXX", -1)
	rewriteXCtoLXXXX := strings.Replace(rewriteXLtoXXXX, "XC", "LXXXX", -1)
	rewriteCDtoCCCC := strings.Replace(rewriteXCtoLXXXX, "CD", "CCCC", -1)
	rewriteDCCCCtoCM := strings.Replace(rewriteCDtoCCCC, "DCCCC", "CM", -1)

	return Numerals{rewriteDCCCCtoCM}
}

// Order of rewrite matters
func additiveSuffixToSubtractivePrefixNormalization(n Numerals) Numerals {
	rewriteDCCCCtoCM := strings.Replace(n.value, "DCCCC", "CM", -1)
	rewriteCCCCtoCD := strings.Replace(rewriteDCCCCtoCM, "CCCC", "CD", -1)
	rewriteLXXXXtoXC := strings.Replace(rewriteCCCCtoCD, "LXXXX", "XC", -1)
	rewriteXXXXtoXL := strings.Replace(rewriteLXXXXtoXC, "XXXX", "XL", -1)
	rewriteVIIIItoIX := strings.Replace(rewriteXXXXtoXL, "VIIII", "IX", -1)
	rewriteIIIItoIV := strings.Replace(rewriteVIIIItoIX, "IIII", "IV", -1)
	return Numerals{rewriteIIIItoIV}
}

// Order of rewrite matters
func additiveSuffixToSingleSuffixNormalization(n Numerals) Numerals {
	rewriteIIIIItoV := strings.Replace(n.value, "IIIII", "V", -1)
	rewriteVVtoX := strings.Replace(rewriteIIIIItoV, "VV", "X", -1)
	rewriteXXXXXtoL := strings.Replace(rewriteVVtoX, "XXXXX", "L", -1)
	rewriteLLtoC := strings.Replace(rewriteXXXXXtoL, "LL", "C", -1)
	rewriteCCCCCtoD := strings.Replace(rewriteLLtoC, "CCCCC", "D", -1)
	rewriteDDtoM := strings.Replace(rewriteCCCCCtoD, "DD", "M", -1)
	return Numerals{rewriteDDtoM}
}
