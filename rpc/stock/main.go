package main

import (
	"log"
	"net"
	"shop/kitex_gen/shop/stock/stockservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	// 服务注册
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")
	svr := stockservice.NewServer(new(StockServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "shop.stock"}))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
