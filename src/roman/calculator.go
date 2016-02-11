package roman

import (
	"fmt"
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

var numeralValueMap = map[rune]int {
	I: 1,
	V: 5,
	X: 10,
	L: 50,
	C: 100,
	D: 500,
	M: 1000,
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
	appendedNumerals :=  Numerals{n1.value + n2.value}
	return additiveSuffixToSubtractivePrefixReduction(appendedNumerals)
}

func additiveSuffixToSubtractivePrefixReduction(n Numerals) Numerals {
	rewriteDCCCCtoCM := strings.Replace(n.value, "DCCCC", "CM", -1)
	rewriteCCCCtoCD := strings.Replace(rewriteDCCCCtoCM, "CCCC", "CD", -1)
	rewriteLXXXXtoXC := strings.Replace(rewriteCCCCtoCD, "LXXXX", "XC", -1)
	rewriteXXXXtoXL := strings.Replace(rewriteLXXXXtoXC, "XXXX", "XL", -1)
	rewriteVIIIItoIX := strings.Replace(rewriteXXXXtoXL, "VIIII", "IX", -1)
	rewriteIIIItoIV := strings.Replace(rewriteVIIIItoIX, "IIII", "IV", -1)
	return Numerals{rewriteIIIItoIV}
}
