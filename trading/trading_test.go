package trading

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuery(t *testing.T) {
	ts := TradeSystem{}
	itemLists := []struct {
		qty        uint
		name       string
		totalValue float64
	}{
		{2, "Silver", 34},
		{4, "Gold", 57800},
		{20, "Iron", 3910},
	}
	for _, item := range itemLists {
		ts.SetItemWithQty(item.name, item.qty, item.totalValue)
	}
	tests := []struct {
		qty  uint
		name string
		want float64
		err  error
	}{
		{4, "Silver", 68, nil},
		{4, "Gold", 57800, nil},
		{4, "Iron", 782, nil},
		{2, "Gold", 28900, nil},
		{4, "Bronze", 0, fmt.Errorf("Item not found Bronze")},
	}
	for _, test := range tests {
		if got, err := ts.Query(test.name, test.qty); got != test.want || !reflect.DeepEqual(err, test.err) {
			t.Errorf("Query(%s, %d)=(%f,%+v), want=(%f,%+v)", test.name, test.qty, got, err, test.want, test.err)
		}
	}
}
