// Code generated by Kitex v0.5.2. DO NOT EDIT.

package echo

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	echo "github.com/phiphi-tan/orbital-api-gateway/kitex_gen/echo"
)

func serviceInfo() *kitex.ServiceInfo {
	return echoServiceInfo
}

var echoServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "Echo"
	handlerType := (*echo.Echo)(nil)
	methods := map[string]kitex.MethodInfo{
		"echo": kitex.NewMethodInfo(echoHandler, newEchoEchoArgs, newEchoEchoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "echo",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func echoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.EchoEchoArgs)
	realResult := result.(*echo.EchoEchoResult)
	success, err := handler.(echo.Echo).Echo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newEchoEchoArgs() interface{} {
	return echo.NewEchoEchoArgs()
}

func newEchoEchoResult() interface{} {
	return echo.NewEchoEchoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Echo(ctx context.Context, req *echo.Request) (r *echo.Response, err error) {
	var _args echo.EchoEchoArgs
	_args.Req = req
	var _result echo.EchoEchoResult
	if err = p.c.Call(ctx, "echo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
