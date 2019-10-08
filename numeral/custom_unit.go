package numeral

import (
	"fmt"
	"strings"
)

type CustomUnit struct {
	units map[string]string
}

// AddUnit Add Unit to dictionary
func (u *CustomUnit) AddUnit(word string, symbol string) {
	u.units[word] = symbol
}

// SetUnits Set Units to dictionary
func (u *CustomUnit) SetUnits(units map[string]string) {
	u.units = units
}

// ToRoman convert customunit to roman
func (u *CustomUnit) ToRoman(s string) (string, error) {
	sTemp := s
	for word, symbol := range u.units {
		sTemp = strings.ReplaceAll(sTemp, word, symbol)
	}
	sTemp = strings.ReplaceAll(sTemp, " ", "")
	_, err := FromRoman(sTemp)
	if err != nil {
		return "", fmt.Errorf("Cannot convert to roman %s", s)
	}
	return sTemp, nil
}

// ToRoman convert customunit to roman
func (u *CustomUnit) ToInt(s string) (int, error) {
	sTemp := s
	for word, symbol := range u.units {
		sTemp = strings.ReplaceAll(sTemp, word, symbol)
	}
	sTemp = strings.ReplaceAll(sTemp, " ", "")
	i, err := FromRoman(sTemp)
	if err != nil {
		return 0, fmt.Errorf("Cannot convert to int %s", s)
	}
	return i, nil
}
