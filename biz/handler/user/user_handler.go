// Code generated by hertz generator.

package user

import (
	"context"
	userService "nodolist/biz/service/user"
	"nodolist/pkg/errno"
	"nodolist/pkg/pack"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	user "nodolist/biz/model/user"
)

// Register .
// @Summary Register
// @Description userRegister
// @Accept json/form
// @Produce json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @router /nodolist/user/register/ [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.RegisterResp)

	userResp, err := userService.NewUserService(ctx).Register(&req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		c.JSON(consts.StatusOK, resp)
		return
	}
	resp.UserID = userResp.ID
	c.JSON(consts.StatusOK, resp)
}

// Login .
// @Summary Login
// @Description login to get your auth token
// @Accept json/form
// @Produce json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @router /nodolist/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	resp := new(user.LoginResp)

	resp.Base = pack.BuildBaseResp(nil)
	//hertz jwt(mw)
	v1, _ := c.Get("user")
	resp.User = userService.BuildUserResp(v1)
	//hertz jwt(mw)
	v2, _ := c.Get("token")
	resp.Token = v2.(string)

	c.JSON(consts.StatusOK, resp)
}

// Info .
// @Summary Information
// @Description show user's info
// @Accept json/form
// @Produce json
// @Param Authorization header string true "token"
// @router /nodolist/user/ [GET]
func Info(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.InfoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.InfoResp)

	v, ok := c.Get("current_user_id")
	if !ok {
		err = errno.ParamError
	}
	id, _ := v.(int64)
	UserResp, err := userService.NewUserService(ctx).Info(id)
	//hertz jwt(mw)

	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp.User = userService.BuildUserResp(UserResp)
	c.JSON(consts.StatusOK, resp)
}