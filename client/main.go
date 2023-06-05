package main

import (
	"context"
	"log"

	"server1/kitex_gen/api"
	"server1/kitex_gen/api/server1"

	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
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
	client, err := server1.NewClient("server1", client.WithResolver(resolver.NewNacosResolver(cli)))
	if err != nil {
		log.Fatal(err)
	}

	addReq := &api.AddRequest{First: 1, Second: 2}
	addResp, err := client.Add(context.Background(), addReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(addResp)

	subReq := &api.SubtractRequest{First: 10, Second: 4}
	subResp, err := client.Subtract(context.Background(), subReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(subResp)

	mulReq := &api.MultiplyRequest{First: 5, Second: 5}
	mulResp, err := client.Multiply(context.Background(), mulReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(mulResp)
}
