package trading

import "fmt"

//TradeSystem implementing trading system
type TradeSystem struct {
	priceList map[string]float64
}

//SetItem add/update item to priceList
func (t *TradeSystem) SetItem(itemName string, value float64) {
	if t.priceList == nil {
		t.priceList = map[string]float64{}
	}
	t.priceList[itemName] = value
}

//SetItemWithQty add/update item to priceList
func (t *TradeSystem) SetItemWithQty(itemName string, qty uint, totalValue float64) {
	t.SetItem(itemName, totalValue/float64(qty))
}

//Query get price query for itemName for qty
func (t *TradeSystem) Query(itemName string, qty uint) (float64, error) {
	if _, ok := t.priceList[itemName]; !ok {
		return 0, fmt.Errorf("Item not found %s", itemName)
	}
	return t.priceList[itemName] * float64(qty), nil
}
