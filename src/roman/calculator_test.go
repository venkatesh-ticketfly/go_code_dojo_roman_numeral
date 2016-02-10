package roman

import "testing"

func TestNewRomanNumaralsAcceptValidInput(t *testing.T) {
	tests := []struct {
		input string
		expected Numerals
	} {
		{"MDCLXVI", Numerals{"MDCLXVI"}},
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
			t.Log(err)
			return
		}
		t.Fatal("Expected an error but was ", actual)
	}
}

func TestRomanNumeralAddition(t *testing.T) {
	tests := []struct {
		num1 Numerals
		num2 Numerals
		expected Numerals
	} {
		{Numerals{"I"}, Numerals{"I"}, Numerals{"II"}},
	}

	for _, test := range tests {
		actual := test.num1.Add(test.num2)

		if actual != test.expected {
			t.Error("Expected ", test.expected, " but was ", actual)
		}
	}
}
