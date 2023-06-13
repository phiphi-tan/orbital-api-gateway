package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	timesvc "github.com/phiphi-tan/orbital-api-gateway/kitex_gen/timesvc/timesvc"
)

func main() {

	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		log.Println("Nacos Registry error:", err)
	}

	svr := timesvc.NewServer(new(TimeSvcImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "timesvc"}),
		server.WithRegistry(r),
		server.WithServiceAddr(&net.TCPAddr{Port: 8003}),
	)

	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}

}
