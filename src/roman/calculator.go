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

type Numerals interface {
	Value() string
	Add(numerals Numerals) Numerals
}

type numerals string

func NewNumerals(numeral string) (Numerals, error) {
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

func (n numerals) Value() string {
	return string(n)
}

func (n numerals) Add(another Numerals) Numerals {
	return numerals(n.Value() + another.Value())
}
