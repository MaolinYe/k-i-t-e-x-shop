package main

import (
	"log"
	stock "shop/kitex_gen/shop/stock/stockservice"
)

func main() {
	svr := stock.NewServer(new(StockServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
