package main

import (
	"log"
	"shop/kitex_gen/shop/item/itemservice"
	"shop/kitex_gen/shop/stock/stockservice"

	"github.com/cloudwego/kitex/client"
)

func main() {
    itemServiceImpl := new(ItemServiceImpl)
    stockCli, err := stockservice.NewClient("shop.item", client.WithHostPorts("0.0.0.0:8890"))
    if err != nil {
       log.Fatal(err)
    }
    itemServiceImpl.stockCli = stockCli

    svr := itemservice.NewServer(itemServiceImpl)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
