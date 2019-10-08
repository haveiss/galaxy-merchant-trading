package numeral

import "testing"

func TestToRoman(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{1, "I"},
		{5, "V"},
		{10, "X"},
		{50, "L"},
		{100, "C"},
		{500, "D"},
		{1000, "M"},
		{4, "IV"},
		{11, "XI"},
		{14, "XIV"},
		{17, "XVII"},
		{29, "XXIX"},
		{48, "XLVIII"},
		{59, "LIX"},
		{94, "XCIV"},
		{145, "CXLV"},
		{469, "CDLXIX"},
		{755, "DCCLV"},
		{999, "CMXCIX"},
		{1903, "MCMIII"},
		{1989, "MCMLXXXIX"},
	}
	for _, test := range tests {
		if got := ToRoman(test.input); got != test.want {
			t.Errorf("ToRoman(%d)=%s, want=%s", test.input, got, test.want)
		}
	}
}

func TestFromRoman(t *testing.T) {
	tests := []struct {
		input string
		want  int
		err   error
	}{
		{"I", 1, nil},
		{"V", 5, nil},
		{"X", 10, nil},
		{"L", 50, nil},
		{"C", 100, nil},
		{"D", 500, nil},
		{"M", 1000, nil},
		{"IV", 4, nil},
		{"XI", 11, nil},
		{"XIV", 14, nil},
		{"XVII", 17, nil},
		{"XXIX", 29, nil},
		{"XLVIII", 48, nil},
		{"LIX", 59, nil},
		{"XCIV", 94, nil},
		{"CXLV", 145, nil},
		{"CDLXIX", 469, nil},
		{"DCCLV", 755, nil},
		{"CMXCIX", 999, nil},
		{"MCMIII", 1903, nil},
		{"MCMLXXXIX", 1989, nil},
	}

	for _, test := range tests {
		if got, err := FromRoman(test.input); got != test.want || err != test.err {
			t.Errorf("FromRoman(%s)=(%d, %+v), want=(%d, %+v)", test.input, got, err, test.want, test.err)
		}
	}
}
