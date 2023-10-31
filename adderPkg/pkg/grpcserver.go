package pkg

import (
	"context"
	"grpc/adderPkg/adder"
)

type GRPCServer struct {
	adder.UnimplementedAdderServer
}

func (s *GRPCServer) Add(ctx context.Context, req *adder.AddRequest) (*adder.AddResponse, error) {
	return &adder.AddResponse{Res: req.GetX() + req.GetY()}, nil
}
