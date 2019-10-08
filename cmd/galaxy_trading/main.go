package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/haveiss/galaxy-merchant-trading/numeral"
	"github.com/haveiss/galaxy-merchant-trading/parser"
	"github.com/haveiss/galaxy-merchant-trading/trading"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	units := numeral.CustomUnit{}
	tradeSystem := trading.TradeSystem{}

	for scanner.Scan() {
		instructionType, args, err := parser.ParseInstruction(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		switch instructionType {
		case "AddUnit":
			addUnitArgs := args.(parser.AddUnitArgs)
			units.AddUnit(addUnitArgs.Word, addUnitArgs.Symbol)
		case "SetItemWithQty":
			setItemWithQtyArgs := args.(parser.SetItemWithQtyArgs)
			i, err := units.ToInt(setItemWithQtyArgs.Words)

			if err != nil {
				fmt.Println(err)
				continue
			}

			tradeSystem.SetItemWithQty(setItemWithQtyArgs.ItemName, uint(i), setItemWithQtyArgs.TotalValue)

		case "ToInt":
			toIntArgs := args.(parser.ToIntArgs)
			i, err := units.ToInt(toIntArgs.Words)

			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("%s is %d\n", toIntArgs.Words, i)

		case "Query":
			queryArgs := args.(parser.QueryArgs)
			i, err := units.ToInt(queryArgs.Words)

			if err != nil {
				fmt.Println(err)
				continue
			}
			f, err := tradeSystem.Query(queryArgs.ItemName, uint(i))

			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("%s %s is %.0f Credits\n", queryArgs.Words, queryArgs.ItemName, f)

		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
