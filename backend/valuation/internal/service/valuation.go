package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"valuation/internal/biz"

	pb "valuation/api/valuation"
)

type ValuationService struct {
	pb.UnimplementedValuationServer
	vlb *biz.ValuationBiz
}

func NewValuationService(vlb *biz.ValuationBiz) *ValuationService {
	return &ValuationService{
		vlb: vlb,
	}
}

func (s *ValuationService) GetValuation(ctx context.Context, req *pb.GetEstimatePriceReq) (*pb.GetEstimatePriceReply, error) {
	distance, duration, err := s.vlb.GetDrivingInfo(context.Background(), req.Origin, req.Destination)
	log.Info(distance,duration)
	if err!=nil {
		return &pb.GetEstimatePriceReply{
			Origin:      req.Origin,
			Destination: req.Destination,
			Price:       0,
		},err
	}
	price,err:=s.vlb.GetPrice(context.Background(),distance,duration,1,9)
	if err!=nil{
		return &pb.GetEstimatePriceReply{
			Origin:      req.Origin,
			Destination: req.Destination,
			Price:       0,
		},err
	}
	return &pb.GetEstimatePriceReply{
		Origin:      req.Origin,
		Destination: req.Destination,
		Price:       price,
	}, nil
}
