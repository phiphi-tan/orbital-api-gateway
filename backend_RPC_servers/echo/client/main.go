package main

import (
	"context"
	"log"

	"github.com/phiphi-tan/orbital-api-gateway/kitex_gen/echo"
	clientecho "github.com/phiphi-tan/orbital-api-gateway/kitex_gen/echo/echo"

	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
)

func main() {
	r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		panic(err)
	}
	client, err := clientecho.NewClient("echo", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}

	req := &echo.Request{Message: "my request"}
	resp, err := client.Echo(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
