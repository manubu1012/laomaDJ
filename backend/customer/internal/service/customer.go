package service

import (
	"context"
	pb "customer/api/customer"
	"customer/api/verifyCode"
	"customer/internal/biz"
	"customer/internal/data"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/wrr"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/hashicorp/consul/api"
	"log"


	"github.com/go-kratos/kratos/v2/transport/grpc"

	"regexp"

	"time"
)

type CustomerService struct {
	pb.UnimplementedCustomerServer
	Cd *data.CustomerData
	cbz *biz.CustomerBiz
}


func NewCustomerService(cd  *data.CustomerData,cbz *biz.CustomerBiz) *CustomerService {
	return &CustomerService{
		Cd: cd,
		cbz: cbz,
	}
}

func (s *CustomerService) EstimatePrice(ctx context.Context, req *pb.EstimatePriceReq) (*pb.EstimatePriceResp, error) {
	price, err := s.cbz.GetEstimatePrice(req.Origin, req.Destination)
	if err!=nil {
		return &pb.EstimatePriceResp{
			Code:        1,
			Message:     "fail",
			Origin:      req.Origin,
			Destination: req.Destination,
			Price:       0,
		},nil
	}
	return &pb.EstimatePriceResp{
		Code:        0,
		Message:     "success",
		Origin:      req.Origin,
		Destination: req.Destination,
		Price: price,
	}, nil
}
func (s *CustomerService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeReq) (*pb.GetVerifyCodeResp, error) {
	pattern := `^(13\d|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18\d|19[0-35-9])\d{8}$`
	regexpPattern:=regexp.MustCompile(pattern)
	if !regexpPattern.MatchString(req.Telephone){
		return &pb.GetVerifyCodeResp{
			Code: 1,
			Message: "手机号不正确",
		},nil
	}
	//和grpc建立连接
	consulConfig:=api.DefaultConfig()
	consulConfig.Address="localhost:8500"
	c, err := api.NewClient(consulConfig)
	if err != nil {
		panic(err)
	}
	// new reg with consul client
	 dis := consul.New(c)
	selector.SetGlobalSelector(wrr.NewBuilder())
	endPoint:="discovery:///verifyCode"
	conn,err:=grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endPoint),
		grpc.WithDiscovery(dis),
	)
	if err!=nil{
		return &pb.GetVerifyCodeResp{
			Code: 1,
			Message: "验证码服务不可达",
		},nil
	}
	defer func() {conn.Close()}()
	//通过client访问
	client:=verifyCode.NewVerifyCodeClient(conn)
	reply,err:=client.GetVerifyCode(context.Background(),&verifyCode.GetVerifyCodeRequest{
		Length: 6,
		Type: 1,
	})
	if err!=nil{
		return &pb.GetVerifyCodeResp{
			Code: 1,
			Message: "验证码服务不可达",
		},nil
	}
	const life=60
	err=s.Cd.SetVerifyCode(req.Telephone,reply.Code,life)
	if err!=nil{
		return &pb.GetVerifyCodeResp{
			Code: 1,
			Message: "验证码存储失败",
		},nil
	}
	return &pb.GetVerifyCodeResp{
		Code: 0,
		VerifyCode: reply.Code,
		VerifyCodeTime: time.Now().Unix(),
		VerifyCodeLife: life,
	}, nil
}


func (s *CustomerService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	////校验验证码
	code:=s.Cd.GetVerifyCode(req.Telephone)
	log.Printf(code)
	if code==""|| req.VerifyCode!=code{
		log.Printf(req.VerifyCode,code)
		log.Printf(code,req.VerifyCode)
		return &pb.LoginResp{
			Code: 1,
			Message: "验证码错误",
		},nil
	}
	//获取用户
	c, err := s.Cd.GetCustomerByTelephone(req.Telephone)
	if err!=nil {
		return &pb.LoginResp{
			Code: 1,
			Message: "获取用户失败",
		},nil
	}
	//生成token

	token,err:=s.Cd.GenerateTokenAndSave(c,biz.Duration,biz.CustomerSecret)
	if err!=nil{
		return &pb.LoginResp{
			Code: 1,
			Message: "生成token错误",
		},nil
	}
	return &pb.LoginResp{
		Code: 0,
		Message: "success",
		Token: token,
		TokenCreateAt: time.Now().Unix(),
		TokenLife: int64(biz.Duration),
	}, nil
}


func (s *CustomerService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutResp, error) {
	//获取用户id
	claims,_:=jwt.FromContext(ctx)
	m:=claims.(jwtv4.MapClaims)
	id:=m["jti"]
	//删除用户的token
	if err := s.Cd.DeleteToken(id);err!=nil{
		return &pb.LogoutResp{
			Code:    1,
			Message: "退出失败",
		},nil
	}
	return &pb.LogoutResp{
		Code:    0,
		Message: "退出成功",
	}, nil


}