// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v3.13.0
// source: api/driver/driver.proto

package driver

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationDriverGetVerifyCode = "/api.driver.Driver/GetVerifyCode"
const OperationDriverLogin = "/api.driver.Driver/Login"
const OperationDriverLogout = "/api.driver.Driver/Logout"

type DriverHTTPServer interface {
	GetVerifyCode(context.Context, *GetVerifyCodeReq) (*GetVerifyCodeResp, error)
	Login(context.Context, *LoginReq) (*LoginResp, error)
	Logout(context.Context, *LogoutReq) (*LogoutResp, error)
}

func RegisterDriverHTTPServer(s *http.Server, srv DriverHTTPServer) {
	r := s.Route("/")
	r.GET("/driver/get-verify-code/{telephone}", _Driver_GetVerifyCode0_HTTP_Handler(srv))
	r.GET("/driver/logout", _Driver_Logout0_HTTP_Handler(srv))
	r.POST("/driver/login", _Driver_Login0_HTTP_Handler(srv))
}

func _Driver_GetVerifyCode0_HTTP_Handler(srv DriverHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetVerifyCodeReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDriverGetVerifyCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetVerifyCode(ctx, req.(*GetVerifyCodeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetVerifyCodeResp)
		return ctx.Result(200, reply)
	}
}

func _Driver_Logout0_HTTP_Handler(srv DriverHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDriverLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutResp)
		return ctx.Result(200, reply)
	}
}

func _Driver_Login0_HTTP_Handler(srv DriverHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDriverLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginResp)
		return ctx.Result(200, reply)
	}
}

type DriverHTTPClient interface {
	GetVerifyCode(ctx context.Context, req *GetVerifyCodeReq, opts ...http.CallOption) (rsp *GetVerifyCodeResp, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginResp, err error)
	Logout(ctx context.Context, req *LogoutReq, opts ...http.CallOption) (rsp *LogoutResp, err error)
}

type DriverHTTPClientImpl struct {
	cc *http.Client
}

func NewDriverHTTPClient(client *http.Client) DriverHTTPClient {
	return &DriverHTTPClientImpl{client}
}

func (c *DriverHTTPClientImpl) GetVerifyCode(ctx context.Context, in *GetVerifyCodeReq, opts ...http.CallOption) (*GetVerifyCodeResp, error) {
	var out GetVerifyCodeResp
	pattern := "/driver/get-verify-code/{telephone}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDriverGetVerifyCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DriverHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginResp, error) {
	var out LoginResp
	pattern := "/driver/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDriverLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DriverHTTPClientImpl) Logout(ctx context.Context, in *LogoutReq, opts ...http.CallOption) (*LogoutResp, error) {
	var out LogoutResp
	pattern := "/driver/logout"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDriverLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
