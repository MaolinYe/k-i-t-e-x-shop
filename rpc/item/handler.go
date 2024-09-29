package main

import (
	"context"
	"shop/kitex_gen/shop/item"
	"shop/kitex_gen/shop/stock"
	"shop/kitex_gen/shop/stock/stockservice"
	"time"

	"github.com/cloudwego/kitex/client/callopt"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct{
	stockCli stockservice.Client
}

// GetItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItem(ctx context.Context, req *item.GetItemReq) (resp *item.GetItemResp, err error) {
	// TODO: Your code here...
	resp = item.NewGetItemResp()
	resp.Item = item.NewItem()
	resp.Item.Id = req.GetId()
	resp.Item.Title = "Kitex"
	resp.Item.Description = "Kitex is an excellent framework!"
	stockReq:=stock.NewGetItemStockReq()
	stockReq.ItemId=req.GetId()
	stockResp,err:=s.stockCli.GetItemStock(ctx,stockReq,callopt.WithRPCTimeout(3*time.Second))
	if err==nil{
		resp.Item.Stock=stockResp.GetStock()
		resp.BaseResp=stockResp.GetBaseResp()
	}
	return
}
