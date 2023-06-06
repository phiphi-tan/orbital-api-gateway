package main

import (
	"context"
	api "hertz_demo/kitex_gen/api"
)

// Server1Impl implements the last service interface defined in the IDL.
type Server1Impl struct{}

// Add implements the Server1Impl interface.
func (s *Server1Impl) Add(ctx context.Context, req *api.AddRequest) (resp *api.AddResponse, err error) {
	// TODO: Your code here...
	return
}

// Subtract implements the Server1Impl interface.
func (s *Server1Impl) Subtract(ctx context.Context, req *api.SubtractRequest) (resp *api.SubtractResponse, err error) {
	// TODO: Your code here...
	return
}

// Multiply implements the Server1Impl interface.
func (s *Server1Impl) Multiply(ctx context.Context, req *api.MultiplyRequest) (resp *api.MultiplyResponse, err error) {
	// TODO: Your code here...
	return
}
