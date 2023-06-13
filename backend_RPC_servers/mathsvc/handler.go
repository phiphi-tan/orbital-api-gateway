package main

import (
	"context"
	"fmt"

	mathsvc "github.com/phiphi-tan/orbital-api-gateway/kitex_gen/mathsvc"
)

// MathSvcImpl implements the last service interface defined in the IDL.
type MathSvcImpl struct{}

// Add implements the MathSvcImpl interface.
func (s *MathSvcImpl) Add(ctx context.Context, req *mathsvc.AddRequest) (resp *mathsvc.AddResponse, err error) {
	fmt.Println("Add Called")
	resp = &mathsvc.AddResponse{Addresult: req.First + req.Second}
	return
}

// Subtract implements the MathSvcImpl interface.
func (s *MathSvcImpl) Subtract(ctx context.Context, req *mathsvc.SubtractRequest) (resp *mathsvc.SubtractResponse, err error) {
	fmt.Println("Subtract Called")
	resp = &mathsvc.SubtractResponse{Subtractresult: req.First - req.Second}
	return
}

// Multiply implements the MathSvcImpl interface.
func (s *MathSvcImpl) Multiply(ctx context.Context, req *mathsvc.MultiplyRequest) (resp *mathsvc.MultiplyResponse, err error) {
	fmt.Println("Multiply Called")
	resp = &mathsvc.MultiplyResponse{Multiplyresult: req.First * req.Second}
	return
}
