package main

import (
	"context"
	"shop/kitex_gen/shop/stock"
	"shop/kitex_gen/shop/base"
)

// StockServiceImpl implements the last service interface defined in the IDL.
type StockServiceImpl struct{}

// GetItemStock implements the StockServiceImpl interface.
func (s *StockServiceImpl) GetItemStock(ctx context.Context, req *stock.GetItemStockReq) (resp *stock.GetItemStockResp, err error) {
	// TODO: Your code here...
	resp=stock.NewGetItemStockResp()
	resp.Stock=req.GetItemId()
	resp.BaseResp=base.NewBaseResp()
	resp.BaseResp.Code="666"
	resp.BaseResp.Msg="success"
	return
}
