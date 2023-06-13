package main

import (
	"context"
	"fmt"

	echo "github.com/phiphi-tan/orbital-api-gateway/kitex_gen/echo"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *echo.Request) (resp *echo.Response, err error) {
	fmt.Println("Echo Called")
	resp = &echo.Response{Message: req.Message}
	return
}
