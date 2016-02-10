package roman

import "testing"

func TestNewRomanNumaralsAcceptValidInput(t *testing.T) {
	tests := []struct {
		input string
		expected numerals
	} {
		{"XVI", numerals("XVI")},
	}

	for _, test := range tests {
		actual, err := NewRomanNeumarals(test.input)
		if err != nil {
			t.Fatal("Expected ", test.expected, " but was ", err)
		}
		if actual != test.expected {
			t.Fatal("Expected ", test.expected, " but was ", actual)
		}
	}
}

func TestNewRomanNumaralsAcceptInvalidInput(t *testing.T) {
	tests := []string{"0"}

	for _, input := range tests {
		actual, err := NewRomanNeumarals(input)
		if err != nil {
			return
		}
		t.Fatal("Expected an error but was ", actual)
	}
}
