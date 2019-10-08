package parser

import (
	"fmt"
	"regexp"
	"strconv"
)

//ParseInstruction parse instruction to instruction_type, parameters and error
func ParseInstruction(s string) (string, interface{}, error) {

	patterns := map[string]*regexp.Regexp{
		"ToInt":          regexp.MustCompile("how much is (?P<Words>.*)\\s+?"),
		"Query":          regexp.MustCompile("how many Credits is (?P<Words>.*)\\s+(?P<ItemName>.*)\\s+?"),
		"SetItemWithQty": regexp.MustCompile("(?P<Words>.*) (?P<ItemName>.*) is (?P<TotalValue>[0-9]+) Credits"),
		"AddUnit":        regexp.MustCompile("(?P<Unit>[a-zA-Z]+) is (?P<Symbol>[a-zA-Z]+)"),
	}

	if patterns["ToInt"].MatchString(s) {
		matches := patterns["ToInt"].FindStringSubmatch(s)
		return "ToInt", ToIntArgs{matches[1]}, nil
	}

	if patterns["Query"].MatchString(s) {
		matches := patterns["Query"].FindStringSubmatch(s)
		return "Query", QueryArgs{matches[1], matches[2]}, nil
	}

	if patterns["SetItemWithQty"].MatchString(s) {
		matches := patterns["SetItemWithQty"].FindStringSubmatch(s)
		totalValue, _ := strconv.ParseFloat(matches[3], 64)
		return "SetItemWithQty", SetItemWithQtyArgs{matches[1], matches[2], totalValue}, nil
	}

	if patterns["AddUnit"].MatchString(s) {
		matches := patterns["AddUnit"].FindStringSubmatch(s)
		return "AddUnit", AddUnitArgs{matches[1], matches[2]}, nil

	}

	return "", nil, fmt.Errorf("I have no idea what you are talking about")
}

//AddUnitArgs ParseInstruction emit AddUnitArgs when instruction matched pattern
type AddUnitArgs struct {
	Word   string
	Symbol string
}

//SetItemWithQtyArgs ParseInstruction emit SetItemWithQtyArgs when instruction matched pattern
type SetItemWithQtyArgs struct {
	Words      string
	ItemName   string
	TotalValue float64
}

//ToIntArgs ParseInstruction emit ToIntArgs when instruction matched pattern
type ToIntArgs struct {
	Words string
}

//QueryArgs ParseInstruction emit QueryArgs when instruction matched pattern
type QueryArgs struct {
	Words    string
	ItemName string
}
