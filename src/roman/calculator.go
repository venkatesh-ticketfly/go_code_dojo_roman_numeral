package roman

import "fmt"

const (
	I = rune('I')
	V = rune('V')
	X = rune('X')
	L = rune('L')
	C = rune('C')
	D = rune('D')
	M = rune('M')
)

type numerals string

func NewRomanNeumarals(numeral string) (numerals, error) {
	nmls := make([]rune, len(numeral))
	for idx, char := range numeral {
		switch char {
		case I, V, X, L, C, D, M:
			nmls[idx] = char
		default:
			return numerals(""), fmt.Errorf("%v is not a valid roman numeral", numeral)
		}
	}
	return numerals(nmls), nil
}
