package numeral

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCustomUnitToRoman(t *testing.T) {
	units := map[string]string{
		"glob": "I",
		"prok": "V",
		"pish": "X",
		"tegj": "L",
	}
	tests := []struct {
		units map[string]string
		input string
		want  string
		err   error
	}{
		{units, "glob glob", "II", nil},
		{units, "glob prok", "IV", nil},
		{units, "pish glob prok", "XIV", nil},
		{units, "pish pish", "XX", nil},
		{units, "pish tegj glob glob", "XLII", nil},
		{units, "pish hello prok", "", fmt.Errorf("Cannot convert to roman pish hello prok")},
	}
	for _, test := range tests {
		c := CustomUnit{}
		c.SetUnits(test.units)
		if got, err := c.ToRoman(test.input); got != test.want || !reflect.DeepEqual(err, test.err) {
			t.Errorf("CustomUnit.ToRoman(%s)=(%s,%+v), want=(%s,%+v)", test.input, got, err, test.want, test.err)
		}
	}
}

func TestCustomUnitToInt(t *testing.T) {
	units := map[string]string{
		"glob": "I",
		"prok": "V",
		"pish": "X",
		"tegj": "L",
	}
	tests := []struct {
		units map[string]string
		input string
		want  int
		err   error
	}{
		{units, "glob glob", 2, nil},
		{units, "glob prok", 4, nil},
		{units, "pish glob prok", 14, nil},
		{units, "pish pish", 20, nil},
		{units, "pish tegj glob glob", 42, nil},
		{units, "pish hello prok", 0, fmt.Errorf("Cannot convert to int pish hello prok")},
	}
	for _, test := range tests {
		c := CustomUnit{}
		c.SetUnits(test.units)
		if got, err := c.ToInt(test.input); got != test.want || !reflect.DeepEqual(err, test.err) {
			t.Errorf("CustomUnit.ToInt(%s)=(%d,%+v), want=(%d,%+v)", test.input, got, err, test.want, test.err)
		}
	}
}
