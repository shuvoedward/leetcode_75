package leetcode

type RomansNumeral struct {
	Value  int
	Symbol string
}

var romanNumerals = []RomansNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func IntToRoman(num int) string {
	result := ""

	for _, numeral := range romanNumerals {
		for num >= numeral.Value {
			result += numeral.Symbol
			num -= numeral.Value
		}

	}
	return result
}
