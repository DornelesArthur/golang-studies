package main

import (
	"fmt"
	"golang_studies/exercises/udemy/2-market_list/model"
	"time"
)

func main() {
	itens := []model.Product{
		{Name: "Condensed Milk", Value: 8.25, Quantity: 2, Unit: "395g"},
		{Name: "Mayonnaise", Value: 19.99, Quantity: 1, Unit: "390g"},
		{Name: "Chocolate Powder", Value: 12.99, Quantity: 1, Unit: "600g", Brand: "Nescau"},
		{Name: "Bread", Value: 12.99, Quantity: 1, Unit: "400g"},
		{Name: "Lager Beer", Value: 3.75, Quantity: 8, Unit: "350ml", Brand: "Brahma"}}

	marketlist, err := model.NewMarketList("Generic SuperMarket", time.Now().AddDate(0, 0, 5), itens)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(marketlist)
	}

}
