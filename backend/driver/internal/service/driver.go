package service

import (
	"context"
	"driver/internal/biz"
	"time"

	pb "driver/api/driver"
)

type DriverService struct {
	pb.UnimplementedDriverServer
	bz *biz.DriverBiz
}

func NewDriverService(bz *biz.DriverBiz) *DriverService {
	return &DriverService{
		bz:bz,
	}
}

func (s *DriverService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeReq) (*pb.GetVerifyCodeResp, error) {
	code, err := s.bz.GetVerifyCode(ctx, req.Telephone)
	if err!=nil{
		return &pb.GetVerifyCodeResp{
			Code:           1,
			Message:        "fail",
			VerifyCode:     "",
			VerifyCodeTime: 0,
			VerifyCodeLife: 0,
		},nil
	}
	return &pb.GetVerifyCodeResp{
		Code:           0,
		Message:        "success",
		VerifyCode:     code,
		VerifyCodeTime: time.Now().Unix(),
		VerifyCodeLife: 60,
	}, nil
}
func (s *DriverService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutResp, error) {
	return &pb.LogoutResp{}, nil
}
func (s *DriverService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	return &pb.LoginResp{}, nil
}
