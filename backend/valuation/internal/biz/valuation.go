package biz

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"

	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/wrr"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	"gorm.io/gorm"
	"strconv"
	"valuation/api/mapService"
)

type  PriceRule struct {
	gorm.Model
	PriceRuleWork
}
type  PriceRuleWork struct {
	CityID      uint  `json:"city_id,omitempty"`
	StartFee    int64 `json:"start_fee,omitempty"`
	DistanceFee int64 `json:"distance_fee,omitempty"`
	DurationFee int64 `json:"duration_fee,omitempty"`
	StartAt     int   `json:"start_at,omitempty"`
	EndAt       int   `json:"end_at,omitempty"`
}

type PriceRuleInterface interface {
	GetPriceRule(cityId uint, curr int)(*PriceRule,error)

}

type ValuationBiz struct {
	pri PriceRuleInterface
}

func NewValuationBiz(pri PriceRuleInterface)*ValuationBiz {
	return &ValuationBiz{
		pri:pri,
	}
}
func (vlb *ValuationBiz)GetPrice(ctx context.Context,distance,duration string, cityId uint,curr int)(int64,error){

	pr,err:=vlb.pri.GetPriceRule(cityId,curr)
	if err!=nil{
		return 0, err
	}
	distanceInt64,err:= strconv.ParseInt(distance,10,64)
	if err!=nil{
		return 0, err
	}
	durationInt64,err:= strconv.ParseInt(duration,10,64)
	if err!=nil{
		return 0, err
	}
	var startDistance int64 = 5
	distanceInt64=distanceInt64/1000
	durationInt64=durationInt64/60
	log.Info(distanceInt64,durationInt64)
	fee:=pr.StartFee+pr.DurationFee*durationInt64+pr.DistanceFee*(distanceInt64-startDistance)
	return fee,nil

}
func (*ValuationBiz)GetDrivingInfo(ctx context.Context,origin,destination string) (string,string,error){
	consulConfig:=api.DefaultConfig()
	consulConfig.Address="localhost:8500"
	c, err := api.NewClient(consulConfig)
	if err != nil {
		panic(err)
	}
	// new reg with consul client
	dis := consul.New(c)
	selector.SetGlobalSelector(wrr.NewBuilder())
	endPoint:="discovery:///map"
	conn,err:=grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endPoint),
		grpc.WithDiscovery(dis),
		grpc.WithMiddleware(
				tracing.Client(),
			),
	)
	if err!=nil{
		return "", "", err
	}
	defer func() {conn.Close()}()
	//通过client访问
	client:=mapService.NewMapServiceClient(conn)
	reply,err:=client.GetMapService(context.Background(),&mapService.GetDrivingInfoReq{
		Origin:      origin,
		Destination: destination,
	})
	if err!=nil{
		 return "", "", err
	}
	return reply.Distance,reply.Duration,nil
}

