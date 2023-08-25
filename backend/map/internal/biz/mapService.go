package biz

import (
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"net/http"
)

type MapServiceBiz struct {
	log  *log.Helper
}

func NewMapServiceBiz(logger log.Logger) *MapServiceBiz {
	return &MapServiceBiz{log:log.NewHelper(logger)}
}

func (m *MapServiceBiz)GetDrivingInfo(origin,destination string) (string,string,error) {
	api:="https://restapi.amap.com/v3/direction/driving"
	key:="22c25418d85beff69d8980c489d3df34"
	params:=fmt.Sprintf("key=%s&origin=%s&destination=%s&extensions=base",key,origin,destination)
	url:=api+"?"+params
	resp, err := http.Get(url)
	if err!=nil{
		log.Info(err)
		return "", "", err
	}
	defer func() {resp.Body.Close()}()
	body,err:=io.ReadAll(resp.Body)
	if err!=nil{
		log.Info(err)
		return "", "", err
	}
	dd:=&DirectionDrivingResp{}
	err = json.Unmarshal(body, dd)
	if err!=nil{
		log.Info(err)
		return "", "", err
	}
	if dd.Info=="OK"{
		return dd.Route.Paths[0].Distance,dd.Route.Paths[0].Duration,nil
	}
	return "","",errors.New(1,"解析错误","获取路线失败")
}

type  DirectionDrivingResp struct {
	Status   string `json:"status,omitempty"`
	Info     string `json:"info,omitempty"`
	InfoCode string `json:"infocode,omitempty"`
	Count    string `json:"count,omitempty"`
	Route    struct {
		Origin      string `json:"origin,omitempty"`
		Destination string `json:"destination,omitempty"`
		Paths       []Path `json:"paths,omitempty"`
	} `json:"route"`
}
type Path struct {
	Distance string `json:"distance,omitempty"`
	Duration string `json:"duration,omitempty"`
}