package parser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {

	type Output struct {
		instructionType string
		params          interface{}
		err             error
	}

	tests := []struct {
		input string
		want  Output
	}{
		{
			"glob is I",
			Output{"AddUnit", AddUnitArgs{"glob", "I"}, nil},
		},
		{
			"glob glob Silver is 34 Credits",
			Output{"SetItemWithQty", SetItemWithQtyArgs{"glob glob", "Silver", 34}, nil},
		},
		{
			"how much is pish tegj glob glob ?",
			Output{"ToInt", ToIntArgs{"pish tegj glob glob"}, nil},
		},
		{
			"how many Credits is glob prok Silver ?",
			Output{"Query", QueryArgs{"glob prok", "Silver"}, nil},
		},
		{
			"how much wood could a woodchuck chuck if a woodchuck could chuck wood ?",
			Output{"", nil, fmt.Errorf("I have no idea what you are talking about")},
		},
	}
	for _, test := range tests {
		instructionType, params, err := ParseInstruction(test.input)
		got := Output{instructionType, params, err}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("ParseInstruction(%s)=%+v, want=%+v", test.input, got, test.want)
		}
	}
}
