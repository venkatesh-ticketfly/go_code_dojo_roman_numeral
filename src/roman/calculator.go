package roman

import (
	"fmt"
	"sort"
	"strings"
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
	subtractiveToAdditiveSuffix := subtractivePrefixToAdditiveSuffixRules.normalize(n)
	return normalizeAdditiveSuffix(subtractiveToAdditiveSuffix)
}

// Order of normalization matters
func normalizeAdditiveSuffix(n Numerals) Numerals {
	sortedNumerals := sortNumerals(n)
	higherNumeralNormalized :=  additiveSuffixToSingleSuffixRules.normalize(sortedNumerals)
	return additiveSuffixToSubtractivePrefixRules.normalize(higherNumeralNormalized)
}

type sortableNumerals []rune

func (s sortableNumerals) Len() int {
	return len(s)
}

func (s sortableNumerals) Less(i, j int) bool {
	return orderedNumeral[s[j]] < orderedNumeral[s[i]]
}

func (s sortableNumerals) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortNumerals(n Numerals) Numerals {
	nRunes := sortableNumerals(n.value)
	sort.Sort(nRunes)
	return Numerals{string(nRunes)}
}

type normalizationRule struct {
	from string
	to   string
}

type normalizationRules []normalizationRule

func (rules normalizationRules) normalize(n Numerals) Numerals {
	convertedNumerals := n.value
	for _, rule := range rules {
		convertedNumerals = strings.Replace(convertedNumerals,
			rule.from, rule.to, -1)
	}

	return Numerals{convertedNumerals}
}

var subtractivePrefixToAdditiveSuffixRules normalizationRules = normalizationRules([]normalizationRule{
	{"IV", "IIII"},
	{"IX", "VIIII"},
	{"XL", "XXXX"},
	{"XC", "LXXXX"},
	{"CD", "CCCC"},
	{"CM", "DCCCC"},
})

var additiveSuffixToSubtractivePrefixRules normalizationRules = normalizationRules([]normalizationRule{
	{"DCCCC", "CM"},
	{"CCCC", "CD"},
	{"LXXXX", "XC"},
	{"XXXX", "XL"},
	{"VIIII", "IX"},
	{"IIII", "IV"},
})

var additiveSuffixToSingleSuffixRules normalizationRules = normalizationRules([]normalizationRule{
	{"IIIII", "V"},
	{"VV", "X"},
	{"XXXXX", "L"},
	{"LL", "C"},
	{"CCCCC", "D"},
	{"DD", "M"},
})
