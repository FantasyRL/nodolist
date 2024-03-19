package jwt

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	"nodolist/biz/model/user"
	userService "nodolist/biz/service/user"
	"nodolist/pkg/errno"
	"nodolist/pkg/pack"
	"time"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	identityKey   = "user_id"
)

func Init() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "nodolist",
		Key:           []byte("nodolist secret key"),
		TokenLookup:   "header:Authorization", //header
		TokenHeadName: "Bearer",               //header
		Timeout:       24 * time.Hour,
		MaxRefresh:    24 * time.Hour,
		IdentityKey:   identityKey,

		// Verify password at login
		//类似于Login Handler
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req user.LoginReq
			err = c.BindAndValidate(&req)
			if err != nil {
				c.String(consts.StatusBadRequest, err.Error())
				return nil, err
			}

			userResp, err := userService.NewUserService(ctx).Login(&req)
			if err != nil {
				return nil, err
			}

			c.Set("user", userResp)

			return userResp.ID, nil
		},
		// Set the payload in the token
		//用于设置登录时为 token 添加自定义负载信息的函数，如果不传入这个参数，
		//则 token 的 payload 部分默认存储 token 的过期时间和创建时间，
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					identityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		// build login response if verify password successfully
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			hlog.CtxInfof(ctx, "Login success ，token is issued clientIP: "+c.ClientIP()) //log
			c.Set("token", token)
		},
		// Verify token and get the id of logged-in user
		//验证用户是否有访问权限，中间件注入返回的就是这个
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			if v, ok := data.(float64); ok {
				current_user_id := int64(v)
				c.Set("current_user_id", current_user_id)
				hlog.CtxInfof(ctx, "Token is verified clientIP: "+c.ClientIP())
				return true
			}
			return false
		},
		// Validation failed, build the message
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			resp := new(user.LoginResp)
			resp.Base = pack.BuildBaseResp(errno.NewErrNo(int64(code), message))
			c.JSON(consts.StatusOK, resp.Base)
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			resp := pack.BuildBaseResp(e)
			return resp.Msg
		},
	})
}
