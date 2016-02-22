package roman

import "testing"

func TestAdditiveSuffixToSubtractivePrefix(t *testing.T) {
	tests := []struct {
		input    Numerals
		expected Numerals
	}{
		{Numerals{"IIII"}, Numerals{"IV"}},
		{Numerals{"VIIII"}, Numerals{"IX"}},
		{Numerals{"XXXX"}, Numerals{"XL"}},
		{Numerals{"LXXXX"}, Numerals{"XC"}},
		{Numerals{"CCCC"}, Numerals{"CD"}},
		{Numerals{"DCCCC"}, Numerals{"CM"}},

		// Rewrite order verification
	}

	for _, test := range tests {
		actual := additiveSuffixToSubtractivePrefixRules.normalize(test.input)

		if actual != test.expected {
			t.Error("Expected ", test.expected, " but was ", actual)
		}
	}
}

func TestAdditiveSuffixToSingleSuffixNormalization(t *testing.T) {
	tests := []struct {
		input    Numerals
		expected Numerals
	}{
		{Numerals{"IIIII"}, Numerals{"V"}},
		{Numerals{"VV"}, Numerals{"X"}},
		{Numerals{"XXXXX"}, Numerals{"L"}},
		{Numerals{"LL"}, Numerals{"C"}},
		{Numerals{"CCCCC"}, Numerals{"D"}},
		{Numerals{"DD"}, Numerals{"M"}},

		{Numerals{"VIIIII"}, Numerals{"X"}},
		{Numerals{"XXXXVIIIII"}, Numerals{"L"}},
		{Numerals{"LXXXXVIIIII"}, Numerals{"C"}},
		{Numerals{"CCCCLXXXXVIIIII"}, Numerals{"D"}},
		{Numerals{"DCCCCLXXXXVIIIII"}, Numerals{"M"}},
	}

	for _, test := range tests {
		actual := additiveSuffixToSingleSuffixRules.normalize(test.input)

		if actual != test.expected {
			t.Error("Expected ", test.expected, " but was ", actual)
		}
	}
}

func TestSubtractivePrefixToAdditiveSuffixNormalization(t *testing.T) {
	tests := []struct {
		input    Numerals
		expected Numerals
	}{
		{Numerals{"IV"}, Numerals{"IIII"}},
		{Numerals{"IX"}, Numerals{"VIIII"}},
		{Numerals{"XL"}, Numerals{"XXXX"}},
		{Numerals{"XC"}, Numerals{"LXXXX"}},
		{Numerals{"CD"}, Numerals{"CCCC"}},
		{Numerals{"CM"}, Numerals{"DCCCC"}},
	}

	for _, test := range tests {
		actual := subtractivePrefixToAdditiveSuffixRules.normalize(test.input)

		if actual != test.expected {
			t.Error("Expected ", test.expected, " but was ", actual)
		}
	}
}
