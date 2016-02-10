package roman

import "fmt"

const (
	I = rune('I')
	V = rune('V')
	X = rune('X')
)

type numerals string

func NewRomanNeumarals(numeral string) (numerals, error) {
	nmls := make([]rune, len(numeral))
	for idx, char := range numeral {
		switch char {
		case 'I':
			nmls[idx] = I
		case 'V':
			nmls[idx] = V
		case 'X':
			nmls[idx] = X
		default:
			return numerals(""), fmt.Errorf("%v is not a valid roman numeral", numeral)
		}
	}
	return numerals(nmls), nil
}
