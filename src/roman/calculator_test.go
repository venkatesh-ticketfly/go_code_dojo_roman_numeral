package roman

import "testing"

func TestNewRomanNumaralsAcceptValidInput(t *testing.T) {
	tests := []struct {
		input string
		expected numerals
	} {
		{"MDCLXVI", numerals("MDCLXVI")},
	}

	for _, test := range tests {
		actual, err := NewNumerals(test.input)
		if err != nil {
			t.Fatal("Expected ", test.expected, " but was ", err)
		}
		if actual != test.expected {
			t.Fatal("Expected ", test.expected, " but was ", actual)
		}
	}
}

func TestNewRomanNumarelsAcceptInvalidInput(t *testing.T) {
	tests := []string{"0", "a", "i"}

	for _, input := range tests {
		actual, err := NewNumerals(input)
		if err != nil {
			return
		}
		t.Fatal("Expected an error but was ", actual)
	}
}

func TestRomanNumeralAddition(t *testing.T) {
	num1, _ := NewNumerals("I")
	num2, _ := NewNumerals("I")

	actual := num1.Add(num2)
	expectedResult, _ := NewNumerals("II")

	if actual != expectedResult {
		t.Fatal("Expected ", expectedResult, " but was ", actual)
	}
}
