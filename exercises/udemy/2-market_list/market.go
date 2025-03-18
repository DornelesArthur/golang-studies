package main

import (
	"fmt"
	"golang_studies/exercises/udemy/2-market_list/model"
	"time"
)

func main() {
	marketlist := model.MarketList{
		Created: time.Now(),
		Date:    time.Now().AddDate(0, 0, 5),
		Market:  "Generic SuperMarket",
	}

	marketlist.Products = append(marketlist.Products, model.Product{Name: "Condensed Milk", Value: 8.25, Quantity: 2, Unit: "395g"})
	marketlist.Products = append(marketlist.Products, model.Product{Name: "Mayonnaise", Value: 19.99, Quantity: 1, Unit: "390g"})
	marketlist.Products = append(marketlist.Products, model.Product{Name: "Chocolate Powder", Value: 12.99, Quantity: 1, Unit: "600g", Brand: "Nescau"})
	marketlist.Products = append(marketlist.Products, model.Product{Name: "Bread", Value: 12.99, Quantity: 1, Unit: "400g"})
	marketlist.Products = append(marketlist.Products, model.Product{Name: "Lager Beer", Value: 3.75, Quantity: 8, Unit: "350ml", Brand: "Brahma"})
	fmt.Println(marketlist)

}
