package calculator

import (
	"context"

	"github.com/fatih/twirpdemo/internal/rpc"
)

// CalculatorService satisfies the Twirp Calculator Interface
type CalculatorService struct{}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

func (c *CalculatorService) Add(ctx context.Context, req *rpc.AddReq) (*rpc.AddResp, error) {
	return &rpc.AddResp{
		Result: req.A + req.B,
	}, nil
}

func (c *CalculatorService) Mul(ctx context.Context, req *rpc.MulReq) (*rpc.MulResp, error) {
	return &rpc.MulResp{
		Result: req.A * req.B,
	}, nil
}
