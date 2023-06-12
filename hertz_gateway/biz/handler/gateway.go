package handler

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/phiphi-tan/orbital-api-gateway/hertz_gateway/biz/errors"
	"github.com/phiphi-tan/orbital-api-gateway/hertz_gateway/biz/types"
	"github.com/phiphi-tan/orbital-api-gateway/kitex_gen/common"
)

type requiredParams struct {
	Method    string `form:"method,required" json:"method"`
	BizParams string `form:"biz_params,required" json:"biz_params"`
}

var SvcMap = make(map[string]genericclient.Client)

// Handler for services
func Gateway(ctx context.Context, c *app.RequestContext) {

	svcName := c.Param("svc")

	//Select Client based on svcName
	cli, ok := SvcMap[svcName]
	if !ok {
		c.JSON(http.StatusOK, errors.New(common.Err_BadRequest))
		return
	}

	//Parameter validation
	var params requiredParams
	if err := c.BindAndValidate(&params); err != nil {
		hlog.Error(err)
		c.JSON(http.StatusOK, errors.New(common.Err_ServerMethodNotFound))
		return
	}

	//Generating HTTP Reqeust
	req, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer([]byte(params.BizParams)))
	if err != nil {
		hlog.Warnf("new http request failed: %v", err)
		c.JSON(http.StatusOK, errors.New(common.Err_RequestServerFail))
		return
	}
	req.URL.Path = fmt.Sprintf("/%s/%s", svcName, params.Method)

	//HTTP to Thrift Mapping
	customReq, err := generic.FromHTTPRequest(req)
	if err != nil {
		hlog.Errorf("convert request failed: %v", err)
		c.JSON(http.StatusOK, errors.New(common.Err_ServerHandleFail))
		return
	}

	//Generic Call
	fmt.Println("Attempting Generic Call")
	resp, err := cli.GenericCall(ctx, "", customReq)
	respMap := make(map[string]interface{})
	//Generic call error handling
	if err != nil {
		hlog.Errorf("GenericCall err:%v", err)
		bizErr, ok := kerrors.FromBizStatusError(err)
		if !ok {
			c.JSON(http.StatusOK, errors.New(common.Err_ServerHandleFail))
			return
		}
		respMap[types.ResponseErrCode] = bizErr.BizStatusCode()
		respMap[types.ResponseErrMessage] = bizErr.BizMessage()
		c.JSON(http.StatusOK, respMap)
		return
	}

	//Reponse Error Handling
	realResp, ok := resp.(*generic.HTTPResponse)
	if !ok {
		c.JSON(http.StatusOK, errors.New(common.Err_ServerHandleFail))
		return
	}

	//Update RequestContext
	realResp.Body[types.ResponseErrCode] = 0
	realResp.Body[types.ResponseErrMessage] = "ok"
	c.JSON(http.StatusOK, realResp.Body)
}
