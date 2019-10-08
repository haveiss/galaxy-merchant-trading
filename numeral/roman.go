package numeral

import "fmt"

var romanUnits = []struct {
	symbol string
	value  int
}{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

// ToRoman Convert Arabic Number to Roman
func ToRoman(i int) string {
	roman := ""
	for _, unit := range romanUnits {
		for i >= unit.value {
			roman = roman + unit.symbol
			i = i - unit.value
		}
	}
	return roman
}

// FromRoman Convert Roman Number to Arabic
func FromRoman(s string) (int, error) {
	return 0, fmt.Errorf("Not implemented yet")
}
