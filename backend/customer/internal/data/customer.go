package data

import (
	"context"

	"customer/internal/biz"
	"database/sql"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomerData struct {
	data *Data
}


func NewCustomerData(data *Data) *CustomerData {
	return &CustomerData{data: data}
}


func (cd *CustomerData)SetVerifyCode(telephone,code string,life int64) error {
	status:=cd.data.Rdb.Set(context.Background(),"CVC"+telephone,code,time.Duration(life)*time.Second)
	if _,err:=status.Result();err!=nil{
		return err
	}
	return nil
}
func (cd *CustomerData)GetVerifyCode(telephone string) string {
	val,_:=cd.data.Rdb.Get(context.Background(),"CVC"+telephone).Result()
	return val
}

func (cd *CustomerData)GetCustomerByTelephone(telephone string) (*biz.Customer,error) {
	customer:=&biz.Customer{}
	cd.data.MDB.Where("telephone = ?",telephone).First(customer)
	if customer.ID>0{
		return customer,nil
	}

		//创建用户
		customer.Telephone=telephone
		resultCreate:=cd.data.MDB.Create(customer)
		if resultCreate.Error==nil{
			return customer,nil
		}else {
			return nil,resultCreate.Error
		}


}



func (cd *CustomerData)GenerateTokenAndSave(c *biz.Customer, d time.Duration,secret string) (string,error) {
	claims:=jwt.RegisteredClaims{
		Issuer:    "LaoMaDJ",
		Subject:   "login-authorization",
		Audience:  []string{"users","others"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(d)),
		NotBefore: nil,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        fmt.Sprintf("%d",c.ID),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken,err:=t.SignedString([]byte(secret))
	l:=log.DefaultLogger
	if err!=nil{
		l.Log(log.LevelInfo,err)
		return "",err
	}

	c.Token=signedToken
	c.TokenCreateAt=sql.NullTime{
		Time: time.Now(),
	}
	if result:=cd.data.MDB.Save(c);result.Error!=nil{
		l.Log(log.LevelInfo,result.Error)
		return "",result.Error
	}
	return signedToken,nil

}

func (cd *CustomerData)GetToken(id interface{}) (string,error) {
	customer := &biz.Customer{}
	if result := cd.data.MDB.First(customer, id); result.Error != nil{
		return "", result.Error
	}
	return customer.Token,nil
}
func (cd *CustomerData)DeleteToken(id interface{}) error{
	customer := &biz.Customer{}
	if result := cd.data.MDB.First(customer, id); result.Error != nil{
		return result.Error
	}
	customer.Token=""
	customer.TokenCreateAt=sql.NullTime{}

	if result := cd.data.MDB.Save(customer); result.Error != nil{
		return result.Error
	}
	return nil
}

