package roman

import "strings"

type normalizationRule struct {
	from string
	to   string
}

type normalizationRules []normalizationRule

func (rules normalizationRules) normalize(n Numerals) Numerals {
	convertedNumerals := n.value
	for _, rule := range rules {
		convertedNumerals = strings.Replace(convertedNumerals, rule.from, rule.to, -1)
	}

	return Numerals{convertedNumerals}
}

var subtractivePrefixToAdditiveSuffixRules normalizationRules = normalizationRules([]normalizationRule{
	{"IV", "IIII"},
	{"IX", "VIIII"},
	{"XL", "XXXX"},
	{"XC", "LXXXX"},
	{"CD", "CCCC"},
	{"CM", "DCCCC"},
})

var additiveSuffixToSubtractivePrefixRules normalizationRules = normalizationRules([]normalizationRule{
	{"DCCCC", "CM"},
	{"CCCC", "CD"},
	{"LXXXX", "XC"},
	{"XXXX", "XL"},
	{"VIIII", "IX"},
	{"IIII", "IV"},
})

var additiveSuffixToSingleSuffixRules normalizationRules = normalizationRules([]normalizationRule{
	{"IIIII", "V"},
	{"VV", "X"},
	{"XXXXX", "L"},
	{"LL", "C"},
	{"CCCCC", "D"},
	{"DD", "M"},
})
