package pack

import (
	"errors"
	"nodolist/biz/model/common"
	"nodolist/pkg/errno"
)

func BuildBaseResp(err error) *common.BaseResp {
	if err == nil {
		return ErrToBaseResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return ErrToBaseResp(e)
	}
	_e := errno.ServiceError.WithMessage(err.Error())
	return ErrToBaseResp(_e)
}

func ErrToBaseResp(err errno.ErrNo) *common.BaseResp {
	return &common.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}
