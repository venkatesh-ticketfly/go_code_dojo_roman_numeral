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

func TestAdditiveSuffixToSubtractivePrefix(t *testing.T) {
	tests := []struct {
		input Numerals
		expected Numerals
	} {
		{Numerals{"IIII"}, Numerals{"IV"}},
		{Numerals{"VIIII"}, Numerals{"IX"}},
		{Numerals{"XXXX"}, Numerals{"XL"}},
		{Numerals{"LXXXX"}, Numerals{"XC"}},
		{Numerals{"CCCC"}, Numerals{"CD"}},
		{Numerals{"DCCCC"}, Numerals{"CM"}},
	}

	for _, test := range tests {
		actual := additiveSuffixToSubtractivePrefixReduction(test.input)

		if actual != test.expected {
			t.Error("Expected ", test.expected, " but was ", actual)
		}
	}
}
