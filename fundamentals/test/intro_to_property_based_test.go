/*
	Some companies will ask you to do the Roman Numeral Kata as part of the interview process.
	This chapter will show how you can tackle it with TDD.
	We are going to write a function which converts an Arabic number (numbers 0 to 9) to a Roman Numeral.
*/

package fundamentalstest

import (
	"fmt"
	"strings"
	"testing"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)

	for _, s := range r {
		if symbol == s.Symbol {
			return s.Value
		}
	}

	return 0
}

var allRomanNumerals = RomanNumerals{
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

var (
	testCases = []struct {
		Arabic int
		Roman  string
	}{
		{Arabic: 1, Roman: "I"},
		{Arabic: 2, Roman: "II"},
		{Arabic: 3, Roman: "III"},
		{Arabic: 4, Roman: "IV"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 6, Roman: "VI"},
		{Arabic: 7, Roman: "VII"},
		{Arabic: 8, Roman: "VIII"},
		{Arabic: 9, Roman: "IX"},
		{Arabic: 10, Roman: "X"},
		{Arabic: 14, Roman: "XIV"},
		{Arabic: 18, Roman: "XVIII"},
		{Arabic: 20, Roman: "XX"},
		{Arabic: 39, Roman: "XXXIX"},
		{Arabic: 40, Roman: "XL"},
		{Arabic: 47, Roman: "XLVII"},
		{Arabic: 49, Roman: "XLIX"},
		{Arabic: 50, Roman: "L"},
		{Arabic: 100, Roman: "C"},
		{Arabic: 90, Roman: "XC"},
		{Arabic: 400, Roman: "CD"},
		{Arabic: 500, Roman: "D"},
		{Arabic: 900, Roman: "CM"},
		{Arabic: 1000, Roman: "M"},
		{Arabic: 1984, Roman: "MCMLXXXIV"},
		{Arabic: 3999, Roman: "MMMCMXCIX"},
		{Arabic: 2014, Roman: "MMXIV"},
		{Arabic: 1006, Roman: "MVI"},
		{Arabic: 798, Roman: "DCCXCVIII"},
	}
)

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	total := 0
	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if couldBeSubstractive(i, symbol, roman) {

			value := allRomanNumerals.ValueOf(symbol, roman[i+1])

			if value != 0 {
				total += value
				i++
			} else {
				total += allRomanNumerals.ValueOf(symbol)
			}
		} else {
			total += allRomanNumerals.ValueOf(symbol)
		}
	}

	return total
}

func couldBeSubstractive(index int, currentSymbol uint8, roman string) bool {
	isSubstractiveSymbol := currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'
	return index+1 < len(roman) && isSubstractiveSymbol

}

func TestRomanNumerals(t *testing.T) {
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%d gets converted to %q", tC.Arabic, tC.Roman), func(t *testing.T) {
			got := ConvertToRoman(tC.Arabic)

			if got != tC.Roman {
				t.Error("want", tC.Roman, "got", got)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, tC := range testCases[:20] {
		t.Run(fmt.Sprintf("%q gets converted to %d", tC.Roman, tC.Arabic), func(t *testing.T) {
			got := ConvertToArabic(tC.Roman)

			if got != tC.Arabic {
				t.Error("want", tC.Arabic, "got", got)
			}
		})
	}
}
