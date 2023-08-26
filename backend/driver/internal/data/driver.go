package data

import (
	"context"
	"time"

	"driver/internal/biz"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/wrr"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"

	"driver/api/verifyCode"
)

type DriverData struct {
	data *Data
}

func NewDriverInterface(data *Data) biz.DriverInterface {
	return &DriverData{
		data: data,
	}
}


func (dd *DriverData) GetVerifyCode(ctx context.Context, tel string) (string, error) {
	consulConfig:=api.DefaultConfig()
	consulConfig.Address=dd.data.cs.Consul.Addr
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
		return "", err
	}
	defer func() {conn.Close()}()
	//通过client访问
	client:=verifyCode.NewVerifyCodeClient(conn)
	reply,err:=client.GetVerifyCode(context.Background(),&verifyCode.GetVerifyCodeRequest{
		Length: 6,
		Type: 1,
	})
	if err!=nil{
		return "", err
	}
	status:=dd.data.RDB.Set(ctx,"DVC"+tel,reply.Code,60*time.Second)

	if _,err:=status.Result();err!=nil{
		return "", err
	}
	return reply.Code,nil

}

