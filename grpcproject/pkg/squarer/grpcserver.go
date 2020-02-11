package squarer

import (
	"context"

	"github.com/vlasove/grpcproject/pkg/api"
)

//GRPCServer ...
type GRPCServer struct{}

//Square ...
func (server *GRPCServer) Square(ctx context.Context, req *api.SquareRequest) (*api.SquareResponse, error) {
	return &api.SquareResponse{Answer: req.GetArg() * req.GetArg()}, nil
}
