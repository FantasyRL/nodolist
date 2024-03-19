package userService

import (
	"nodolist/biz/dal"
	"nodolist/biz/model/user"
	"nodolist/pkg/errno"
	"nodolist/pkg/utils/pwd"
)

func (s *UserService) Register(req *user.RegisterReq) (*dal.User, error) {
	if len(req.Username) < 4 /*||len(req.Password)<8*/ {
		return nil, errno.ParamError
	}

	PwdDigest := pwd.SetPassword(req.Password)
	userModel := &dal.User{
		UserName: req.Username,
		Password: PwdDigest,
	}
	return dal.Register(s.ctx, userModel)
}

func (s *UserService) Login(req *user.LoginReq) (*dal.User, error) {
	userModel := &dal.User{
		UserName: req.Username,
		Password: req.Password,
	}
	userResp, err := dal.Login(s.ctx, userModel)
	if err != nil {
		return nil, err
	}
	return userResp, nil
}

func (s *UserService) Info(id int64) (*dal.User, error) {
	return dal.QueryUserByID(id)
}
