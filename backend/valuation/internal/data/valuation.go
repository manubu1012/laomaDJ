package data

import (
	"github.com/go-kratos/kratos/v2/errors"
	"valuation/internal/biz"
)

type PriceRuleData struct {
	data *Data
}

func NewPriceRuleInterface(data *Data) biz.PriceRuleInterface {
	return &PriceRuleData{data: data}
}

func (vld *PriceRuleData)GetPriceRule(cityId uint, curr int)(*biz.PriceRule,error)  {
	pr:=&biz.PriceRule{}
	result := vld.data.MDB.Where("city_id=? and start_at <=? and end_at>?", cityId, curr, curr).First(pr)
	if result.Error!=nil{
		return nil, errors.New(1,"db not found","not found")
	}
	return pr,nil
}