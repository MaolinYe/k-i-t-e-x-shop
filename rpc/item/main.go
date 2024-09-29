package main

import (
	"log"
	"net"
	"shop/kitex_gen/shop/item/itemservice"
	"shop/kitex_gen/shop/stock/stockservice"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	itemServiceImpl := new(ItemServiceImpl)
	// 服务发现
	resolver, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	// rpc调用货存
	stockCli, err := stockservice.NewClient("shop.stock", client.WithResolver(resolver))
	if err != nil {
		log.Fatal(err)
	}
	itemServiceImpl.stockCli = stockCli
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	// 服务注册
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8891")
	svr := itemservice.NewServer(itemServiceImpl,
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{ServiceName: "shop.item"}))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
