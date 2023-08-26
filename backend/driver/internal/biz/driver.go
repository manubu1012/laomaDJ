package biz

import (
	"context"
	"database/sql"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/errors"
	"regexp"
)

type DriverBiz struct {
	di DriverInterface
}

type Driver struct {
	gorm.Model
	DriverWork
}
type DriverWork struct {
	Telephone     string `gorm:"type:varchar(16);uniqueIndex;" json:"telephone,omitempty"`
	Name          sql.NullString `gorm:"type:varchar(255);index;;" json:"name,omitempty"`
	Token         sql.NullString `gorm:"type:varchar(2047);" json:"token,omitempty"`
	Status        sql.NullString `gorm:"type:enum('in','out','listen','stop');" json:"status,omitempty"`
	IdNumber      sql.NullString `gorm:"type:varchar(18);index;" json:"id_number,omitempty"`
	IdImageA      sql.NullString `gorm:"type:varchar(255);" json:"id_image_a,omitempty"`
	LicenseImageA sql.NullString `gorm:"type:varchar(255);" json:"license_image_a,omitempty"`
	LicenseImageB sql.NullString `gorm:"type:varchar(255);" json:"license_image_b,omitempty"`
	DistinctCode  sql.NullString `gorm:"type:varchar(16);index;" json:"distinct_code,omitempty"`
	AuditAt       sql.NullTime `gorm:"index;" json:"audit_at,omitempty"`
	TelephoneBak  sql.NullString `gorm:"type:varchar(16);" json:"telephone_bak,omitempty"`
}
const DriverStatusIn = "in"
const DriverStatusOut = "out"
const DriverStatusListen = "listen"
const DriverStatusStop = "stop"


type DriverInterface interface {
	GetVerifyCode(context.Context,string)(string,error)
}

func NewDriverBiz(di DriverInterface)*DriverBiz  {
	return &DriverBiz{
		di: di,
	}
}

func (bz *DriverBiz)GetVerifyCode(ctx context.Context,tel string)(string,error)  {
	pattern := `^(13\d|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18\d|19[0-35-9])\d{8}$`
	regexpPattern:=regexp.MustCompile(pattern)
	if !regexpPattern.MatchString(tel){
		return "",errors.New(200,"Driver","Phone Number Error")
	}
	verifyCode,err:= bz.di.GetVerifyCode(ctx,tel)
	if err!=nil{
		return "",errors.New(200,"Redis","Redis Found Error")
	}
	return verifyCode,nil
}


