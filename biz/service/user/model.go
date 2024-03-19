package userService

import (
	"context"
	"nodolist/biz/dal"
	"nodolist/biz/model/user"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func BuildUserResp(_user interface{}) *user.User {
	//这里使用了一个及其抽象的断言
	p, _ := (_user).(*dal.User)
	return &user.User{
		ID:        p.ID,
		Name:      p.UserName,
		TodoCount: dal.GetTodoCountByUid(p.ID),
	}
}
