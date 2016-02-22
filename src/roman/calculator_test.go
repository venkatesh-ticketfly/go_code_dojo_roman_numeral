package roman

import (
	"testing"
)

func TestNewRomanNumaralsAcceptValidInput(t *testing.T) {
	tests := []struct {
		input    string
		expected Numerals
	}{
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

func TestNormalizationAdditiveSuffix(t *testing.T) {
	tests := []struct {
		input    Numerals
		expected Numerals
	}{
		{Numerals{"IIIIIIIII"}, Numerals{"IX"}},
		{Numerals{"VVVVVVVV"}, Numerals{"XL"}},
		{Numerals{"XXXXXXXXX"}, Numerals{"XC"}},
		{Numerals{"LLLLLLLL"}, Numerals{"CD"}},
		{Numerals{"CCCCCCCCC"}, Numerals{"CM"}},

		{Numerals{"XXXVIIIII"}, Numerals{"XL"}},

		// Unsorted
		{Numerals{"IIVIIIXXX"}, Numerals{"XL"}},
	}

	for _, test := range tests {
		actual := normalizeAdditiveSuffix(test.input)

		if actual != test.expected {
			t.Error("Expected ", test.expected, " but was ", actual)
		}
	}
}

func TestNormalization(t *testing.T) {
	tests := []struct {
		input    Numerals
		expected Numerals
	}{
		{Numerals{"XXXIVVI"}, Numerals{"XL"}},
	}

	for _, test := range tests {
		actual := normalize(test.input)

		if actual != test.expected {
			t.Error("Expected ", test.expected, " but was ", actual)
		}
	}
}
