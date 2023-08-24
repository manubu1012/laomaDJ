package data

import (
	"customer/internal/biz"
	"customer/internal/conf"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo,NewCustomerData)

// Data .
type Data struct {
	// TODO wrapped database client
	Rdb *redis.Client
	MDB *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	data:=&Data{}
	//redis配置
	URL:=fmt.Sprintf("redis://%s/1?dial_timeout=%d",c.Redis.Addr,1)
	options,_:=redis.ParseURL(URL)
	rdb := redis.NewClient(options)
	data.Rdb=rdb
	//mysql配置
	db, err:= gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	migrate(db)
	if err !=nil{
		data.MDB=nil
		log.Fatalf("mysql database err\n")
	}
	data.MDB=db
	
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
}

func migrate(db  *gorm.DB) error {
	if err := db.AutoMigrate(&biz.Customer{}); err != nil {
		return err
	}
	return nil
}