package romannumerals

import "errors"

var rn = []struct {
	arabic int
	roman  string
}{
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

func ToRomanNumeral(num int) (string, error) {
	var romanNum string
	if num <= 0 || num >= 4000 {
		return "", errors.New("number is too small or too large.")
	}
	for _, value := range rn {
		for num >= value.arabic {
			romanNum += value.roman
			num -= value.arabic
		}
	}
	return romanNum, nil
}
