package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"valuation/internal/biz"
	"valuation/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo,NewPriceRuleInterface)

// Data .
type Data struct {
	// TODO wrapped database client
	MDB *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	data:=&Data{}

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
	if err := db.AutoMigrate(&biz.PriceRule{}); err != nil {
		return err
	}
	rules := []biz.PriceRule{
		{
			Model: gorm.Model{ID: 1},
			PriceRuleWork: biz.PriceRuleWork{
				CityID:      1,
				StartFee:    300,
				DistanceFee: 35,
				DurationFee: 10, // 5m
				StartAt:     7,
				EndAt:       23,
			},
		},
		{
			Model: gorm.Model{ID: 2},
			PriceRuleWork: biz.PriceRuleWork{
				CityID:      1,
				StartFee:    350,
				DistanceFee: 35,
				DurationFee: 10, // 5m
				StartAt:     23,
				EndAt:       24,
			},
		},
		{
			Model: gorm.Model{ID: 3},
			PriceRuleWork: biz.PriceRuleWork{
				CityID:      1,
				StartFee:    400,
				DistanceFee: 35,
				DurationFee: 10, // 5m
				StartAt:     0,
				EndAt:       7,
			},
		},
	}
	db.Clauses(clause.OnConflict{UpdateAll: true}).Create(rules)
	return nil
}

