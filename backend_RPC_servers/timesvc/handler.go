package main

import (
	"context"
	"fmt"
	"time"

	timesvc "github.com/phiphi-tan/orbital-api-gateway/kitex_gen/timesvc"
)

// TimeSvcImpl implements the last service interface defined in the IDL.
type TimeSvcImpl struct{}

// GetTime implements the TimeSvcImpl interface.
func (s *TimeSvcImpl) GetTime(ctx context.Context, req *timesvc.TimeReq) (resp *timesvc.TimeResp, err error) {
	fmt.Println("GetTime Called")
	currtime := time.Now()
	if req.Twentyfourhour {
		//twentyfourformat := fmt.Sprintf("%d:%d:%d", currtime.Hour(), currtime.Minute(), currtime.Second())
		resp = &timesvc.TimeResp{Time: currtime.Format("15:04:05")}
	} else {
		resp = &timesvc.TimeResp{Time: currtime.Format("03:04:05 PM")}
	}
	return
}

// GetDate implements the TimeSvcImpl interface.
func (s *TimeSvcImpl) GetDate(ctx context.Context, req *timesvc.DateReq) (resp *timesvc.DateResp, err error) {
	fmt.Println("GetDate Called")
	currtime := time.Now()
	if req.AmericanFormatting {
		resp = &timesvc.DateResp{Date: currtime.Format("01/02/2006")}
	} else {
		resp = &timesvc.DateResp{Date: currtime.Format("02/01/2006")}
	}
	return
}
