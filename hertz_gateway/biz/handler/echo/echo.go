// Code generated by hertz generator.

package echo

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
<<<<<<< Updated upstream
)
=======
	clients "github.com/phiphi-tan/orbital-api-gateway/hertz_gateway/biz/clients"
	echo "github.com/phiphi-tan/orbital-api-gateway/hertz_gateway/biz/model/echo"
)

// Echo .
// @router /echo/echo [POST]
func Echo(ctx context.Context, c *app.RequestContext) {

	//Binding and Validation of Params
	var err error
	var req echo.Request
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	clients.GenericCall("echo", "echo", ctx, c)
}
>>>>>>> Stashed changes
