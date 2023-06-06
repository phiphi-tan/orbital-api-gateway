package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"hertz_demo/kitex_gen/api"
	"hertz_demo/kitex_gen/api/server1"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	h := server.Default()

	// ping function
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	// hello function
	h.GET("/hello", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "Hello World!"})
	})

	// welcome function
	h.GET("/welcome", func(c context.Context, ctx *app.RequestContext) {
		firstname := ctx.DefaultQuery("firstname", "Guest")
		// shortcut for c.Request.URL.Query().Get("lastname")
		lastname := ctx.Query("lastname")

		// Iterate all queries and store the one with meeting the conditions in favoriteFood
		var favoriteFood []string
		ctx.QueryArgs().VisitAll(func(key, value []byte) {
			if string(key) == "food" {
				favoriteFood = append(favoriteFood, string(value))
			}
		})

		ctx.String(consts.StatusOK, "Hello %s %s, favorite food: %s", firstname, lastname, favoriteFood)
	})

	// testing function
	h.GET("/addingTest", func(c context.Context, ctx *app.RequestContext) {
		num1 := ctx.DefaultQuery("num1", "0")
		number1, _ := strconv.ParseFloat(num1, 64)

		fmt.Println(number1)
		num2 := ctx.Query("num2")
		number2, _ := strconv.ParseFloat(num2, 64)
		fmt.Println(number2)

		//sum := number1 + number2
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
		client, err := server1.NewClient("server1", client.WithResolver(resolver.NewNacosResolver(cli)))
		//Check load balancer mechanics - Is it an object?
		if err != nil {
			log.Fatal(err)
		}

		addReq := &api.AddRequest{First: number1, Second: number2}
		addResp, err := client.Add(context.Background(), addReq)
		if err != nil {
			log.Fatal(err)
		}

		sumStr := strconv.FormatFloat(addResp.Addresult, 'f', -1, 64)

		ctx.JSON(consts.StatusOK, utils.H{"sum": sumStr})
	})

	h.Spin()

}
