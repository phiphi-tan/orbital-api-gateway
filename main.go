package main

import (
	"log"
	"net"
	api "server1/kitex_gen/api/server1"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	// ...
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("192.168.1.3", 8848),
	}

	cc := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "info",
		Username:            "your-name",
		Password:            "your-password",
	}

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", ":9999")
	svr := api.NewServer(new(Server1Impl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "server1"}),
		server.WithRegistry(registry.NewNacosRegistry(cli)),
		server.WithServiceAddr(addr),
	)
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
