package main

import (
	"context"
	api "server1/kitex_gen/api"
)

// Server1Impl implements the last service interface defined in the IDL.
type Server1Impl struct{}

// Add implements the Server1Impl interface.
func (s *Server1Impl) Add(ctx context.Context, req *api.AddRequest) (resp *api.AddResponse, err error) {
	resp = &api.AddResponse{Addresult: req.First + req.Second}
	return
}

// Subtract implements the Server1Impl interface.
func (s *Server1Impl) Subtract(ctx context.Context, req *api.SubtractRequest) (resp *api.SubtractResponse, err error) {
	resp = &api.SubtractResponse{Subtractresult: req.First - req.Second}
	return
}

// Multiply implements the Server1Impl interface.
func (s *Server1Impl) Multiply(ctx context.Context, req *api.MultiplyRequest) (resp *api.MultiplyResponse, err error) {
	resp = &api.MultiplyResponse{Multiplyresult: req.First * req.Second}
	return
}
