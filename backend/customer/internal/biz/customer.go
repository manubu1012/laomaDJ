package biz

import (
	"context"

	"customer/api/valuation"
	"database/sql"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/wrr"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	"gorm.io/gorm"
	"time"
)

type  Customer struct {
	CustomerWork
	CustomerToken
	gorm.Model

}
type CustomerWork struct {
	Telephone string `gorm:"type:varchar(15);uniqueIndex;"json:"telephone"`
	Name string`gorm:"type:varchar(255);",json:"name"`
	Wechat string`gorm:"type:varchar(255);",json:"wechat"`
	Email string`gorm:"type:varchar(255);",json:"email"`
}

type CustomerToken struct {
	Token string `gorm:"varchar(4095);"json:"token"`
	TokenCreateAt sql.NullTime `json:"token_create_at"`
}

const Duration = 30 * 60 * 60 * 24 * time.Second
const CustomerSecret = "my_secret_key"

type CustomerBiz struct {
}

func NewCustomerBiz() *CustomerBiz {
	return &CustomerBiz{}
}

func (*CustomerBiz)GetEstimatePrice(orgin,destination string)(int64,error)  {
	consulConfig:=api.DefaultConfig()
	consulConfig.Address="localhost:8500"
	c, err := api.NewClient(consulConfig)
	if err != nil {
		panic(err)
	}
	// new reg with consul client
	dis := consul.New(c)
	selector.SetGlobalSelector(wrr.NewBuilder())
	endPoint:="discovery:///valuation"
	conn,err:=grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endPoint),
		grpc.WithDiscovery(dis),
	)
	if err!=nil{
		return 0,err
	}
	defer func() {conn.Close()}()
	//通过client访问
	client:=valuation.NewValuationClient(conn)
	reply,err:=client.GetValuation(context.Background(),&valuation.GetEstimatePriceReq{
		Origin: orgin,
		Destination: destination,
	})
	if err!=nil{
	 return 0,err
	}
	return reply.Price,nil
}