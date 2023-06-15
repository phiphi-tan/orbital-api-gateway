// Code generated by hertz generator. DO NOT EDIT.

package timesvc

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	timesvc "github.com/phiphi-tan/orbital-api-gateway/hertz_gateway/biz/handler/timesvc"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_timesvc := root.Group("/timesvc", _timesvcMw()...)
		_timesvc.POST("/date", append(_getdateMw(), timesvc.GetDate)...)
		_timesvc.POST("/time", append(_gettimeMw(), timesvc.GetTime)...)
	}
}
