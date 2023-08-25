package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"map/internal/biz"

	pb "map/api/mapService"
)

type MapServiceService struct {
	pb.UnimplementedMapServiceServer
	msb *biz.MapServiceBiz
}

func NewMapServiceService() *MapServiceService {
	return &MapServiceService{}
}

func (s *MapServiceService) GetMapService(ctx context.Context, req *pb.GetDrivingInfoReq) (*pb.GetDrivingInfoReply, error) {
	distance, duration, err := s.msb.GetDrivingInfo(req.Origin, req.Destination)
	if err!=nil{
		log.Info(err)
		return &pb.GetDrivingInfoReply{
			Origin:      req.Origin,
			Destination: req.Destination,
			Distance:    "",
			Duration:    "",
		},nil
	}
	return &pb.GetDrivingInfoReply{
		Origin:      req.Origin,
		Destination: req.Destination,
		Distance:    distance,
		Duration:    duration,
	}, nil
}
