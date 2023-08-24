package server

import (
	"context"
	"customer/internal/service"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"strings"
)

func customerJWT(customerService *service.CustomerService) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			claims,ok:=jwt.FromContext(ctx)
			if !ok{
				return nil, errors.Unauthorized("UNAUTHORIZED","claims not found")
			}
			m:=claims.(jwtv4.MapClaims)
			id:=m["jti"]
			token,err:=customerService.Cd.GetToken(id)
			if err!=nil{
				return nil, errors.Unauthorized("UNAUTHORIZED","user not found")
			}
			header,_:=transport.FromServerContext(ctx)
			auths := strings.SplitN(header.RequestHeader().Get("Authorization"), " ", 2)
			jwtToken := auths[1]
			if jwtToken!=token{
				return nil, errors.Unauthorized("UNAUTHORIZED","token was updated")
			}
			return handler(ctx,req)
		}
	}
}
