// Code generated by Kitex v0.5.2. DO NOT EDIT.

package mathsvc

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	mathsvc "github.com/phiphi-tan/orbital-api-gateway/kitex_gen/mathsvc"
)

func serviceInfo() *kitex.ServiceInfo {
	return mathSvcServiceInfo
}

var mathSvcServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "MathSvc"
	handlerType := (*mathsvc.MathSvc)(nil)
	methods := map[string]kitex.MethodInfo{
		"add":      kitex.NewMethodInfo(addHandler, newMathSvcAddArgs, newMathSvcAddResult, false),
		"subtract": kitex.NewMethodInfo(subtractHandler, newMathSvcSubtractArgs, newMathSvcSubtractResult, false),
		"multiply": kitex.NewMethodInfo(multiplyHandler, newMathSvcMultiplyArgs, newMathSvcMultiplyResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "mathsvc",
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

func addHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*mathsvc.MathSvcAddArgs)
	realResult := result.(*mathsvc.MathSvcAddResult)
	success, err := handler.(mathsvc.MathSvc).Add(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMathSvcAddArgs() interface{} {
	return mathsvc.NewMathSvcAddArgs()
}

func newMathSvcAddResult() interface{} {
	return mathsvc.NewMathSvcAddResult()
}

func subtractHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*mathsvc.MathSvcSubtractArgs)
	realResult := result.(*mathsvc.MathSvcSubtractResult)
	success, err := handler.(mathsvc.MathSvc).Subtract(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMathSvcSubtractArgs() interface{} {
	return mathsvc.NewMathSvcSubtractArgs()
}

func newMathSvcSubtractResult() interface{} {
	return mathsvc.NewMathSvcSubtractResult()
}

func multiplyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*mathsvc.MathSvcMultiplyArgs)
	realResult := result.(*mathsvc.MathSvcMultiplyResult)
	success, err := handler.(mathsvc.MathSvc).Multiply(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMathSvcMultiplyArgs() interface{} {
	return mathsvc.NewMathSvcMultiplyArgs()
}

func newMathSvcMultiplyResult() interface{} {
	return mathsvc.NewMathSvcMultiplyResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Add(ctx context.Context, req *mathsvc.AddRequest) (r *mathsvc.AddResponse, err error) {
	var _args mathsvc.MathSvcAddArgs
	_args.Req = req
	var _result mathsvc.MathSvcAddResult
	if err = p.c.Call(ctx, "add", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Subtract(ctx context.Context, req *mathsvc.SubtractRequest) (r *mathsvc.SubtractResponse, err error) {
	var _args mathsvc.MathSvcSubtractArgs
	_args.Req = req
	var _result mathsvc.MathSvcSubtractResult
	if err = p.c.Call(ctx, "subtract", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Multiply(ctx context.Context, req *mathsvc.MultiplyRequest) (r *mathsvc.MultiplyResponse, err error) {
	var _args mathsvc.MathSvcMultiplyArgs
	_args.Req = req
	var _result mathsvc.MathSvcMultiplyResult
	if err = p.c.Call(ctx, "multiply", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}