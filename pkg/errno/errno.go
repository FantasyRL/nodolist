package errno

import (
	"errors"
	"fmt"
)

const (
	WebSocketSuccessCode       = 1000
	WebSocketLogoutSuccessCode = iota + 1000
	WebSocketTargetOnlineSuccessCode
	WebSocketErrCode
	WebSocketConnectErrCode
	WebSocketTargetOfflineErrCode

	SuccessCode    = 10000
	ServiceErrCode = iota + 10000 //未知服务错误
	ParamErrCode                  //参数错误
	CharacterBeyondLimitErrCode

	ExistUserErrCode
	NotExistUserErrCode
	AuthFailedErrCode //Authorization错误
	ReadFileErrCode
	UploadFileErrCode
	LikeExistErrCode
	LikeNotExistErrCode
	VideoNotExistErrCode
	CommentIsNotExistErrCode
	ParentCommentIsNotExistErrCode
	FollowExistErrCode
	FollowNotExistErrCode
	FollowMyselfErrCode
	Enable2FAErrCode
	Unable2FAErrCode
	Verify2FAErrCode

	TaskNotExistErrCode
	TaskStatusErrCode
)

const (
	SuccessMsg                    = "Success"
	ServerErrMsg                  = "Service is unable to start successfully"
	ParamErrMsg                   = "Wrong Parameter has been given"
	CharacterBeyondLimitErrMsg    = "the number of character beyond the limit"
	UserAlreadyExistErrMsg        = "User existed"
	UserIsNotExistErrMsg          = "User is not exist"
	PasswordIsNotVerifiedMsg      = "Username or password not verified"
	AuthErrMsg                    = "It is not your account"
	ReadFileErrMsg                = "Error when read file"
	UploadFileErrMsg              = "Upload file error"
	LikeExistErrMsg               = "You have liked this target"
	LikeNotExistErrMsg            = "You don't like this video"
	LikeActionErrMsg              = "Favorite add failed"
	FollowExistErrMsg             = "You have followed"
	FollowNotExistErrMsg          = "You haven't followed"
	FollowMyselfErrMsg            = "You can't follow yourself"
	MessageAddFailedErrMsg        = "message add failed"
	FriendListNoPermissionMsg     = "You can't query his friend list"
	VideoNotExistErrMsg           = "Video is not exist"
	CommentIsNotExistErrMsg       = "Comment is not exist"
	ParentCommentIsNotExistErrMsg = "Parent-comment is not exist"
	Enable2FAErrMsg               = "2fa verification have opened"
	Unable2FAErrMsg               = "2fa verification have closed"
	Verify2FAErrMsg               = "incorrect otp password"

	WebSocketSuccessMsg             = "Connect to server success"
	WebSocketLogoutSuccessMsg       = "logout success"
	WebSocketTargetOnlineSuccessMsg = "target user is online"
	WebSocketConnectErrMsg          = "Connect or upgrade error"
	WebSocketTargetOfflineErrMsg    = "Target user is offline"
	WebSocketErrMsg                 = "Websocket error"

	TaskNotExistErrMsg = "Task not exist"
	TaskStatusErrMsg   = "status error"
)

type ErrNo struct {
	ErrorCode int64
	ErrorMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("%s", e.ErrorMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrorCode: code,
		ErrorMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo { //出现不被定义的错误时
	e.ErrorMsg = msg
	return e
}

var (
	Success                   = NewErrNo(SuccessCode, SuccessMsg)
	ServiceError              = NewErrNo(ServiceErrCode, ServerErrMsg)
	ParamError                = NewErrNo(ParamErrCode, ParamErrMsg)
	CharacterBeyondLimitError = NewErrNo(CharacterBeyondLimitErrCode, CharacterBeyondLimitErrMsg)

	ExistUserError               = NewErrNo(ExistUserErrCode, UserAlreadyExistErrMsg)
	NotExistUserError            = NewErrNo(NotExistUserErrCode, UserIsNotExistErrMsg)
	PwdError                     = NewErrNo(AuthFailedErrCode, PasswordIsNotVerifiedMsg)
	AuthorizationError           = NewErrNo(AuthFailedErrCode, AuthErrMsg)
	UploadFileError              = NewErrNo(UploadFileErrCode, UploadFileErrMsg)
	ReadFileError                = NewErrNo(ReadFileErrCode, ReadFileErrMsg)
	LikeExistError               = NewErrNo(LikeExistErrCode, LikeExistErrMsg)
	LikeNotExistError            = NewErrNo(LikeNotExistErrCode, LikeNotExistErrMsg)
	VideoNotExistError           = NewErrNo(VideoNotExistErrCode, VideoNotExistErrMsg)
	CommentIsNotExistError       = NewErrNo(CommentIsNotExistErrCode, CommentIsNotExistErrMsg)
	ParentCommentIsNotExistError = NewErrNo(ParentCommentIsNotExistErrCode, ParentCommentIsNotExistErrMsg)
	FollowExistError             = NewErrNo(FollowExistErrCode, FollowExistErrMsg)
	FollowNotExistError          = NewErrNo(FollowNotExistErrCode, FollowNotExistErrMsg)
	FollowMyselfError            = NewErrNo(FollowMyselfErrCode, FollowMyselfErrMsg)
	Enable2FAError               = NewErrNo(Enable2FAErrCode, Enable2FAErrMsg)
	Unable2FAError               = NewErrNo(Unable2FAErrCode, Unable2FAErrMsg)
	Verify2FAError               = NewErrNo(Verify2FAErrCode, Verify2FAErrMsg)

	WebSocketSuccess             = NewErrNo(WebSocketSuccessCode, WebSocketSuccessMsg)
	WebSocketLogoutSuccess       = NewErrNo(WebSocketLogoutSuccessCode, WebSocketLogoutSuccessMsg)
	WebSocketTargetOnlineSuccess = NewErrNo(WebSocketTargetOnlineSuccessCode, WebSocketTargetOnlineSuccessMsg)
	WebSocketConnectError        = NewErrNo(WebSocketConnectErrCode, WebSocketConnectErrMsg)
	WebSocketTargetOfflineError  = NewErrNo(WebSocketTargetOfflineErrCode, WebSocketTargetOfflineErrMsg)
	WebSocketError               = NewErrNo(WebSocketErrCode, WebSocketErrMsg)

	TaskNotExistError = NewErrNo(TaskNotExistErrCode, TaskNotExistErrMsg)
	TaskStatusError   = NewErrNo(TaskStatusErrCode, TaskStatusErrMsg)
)

// ConvertErr convert error to ErrNo
// in Default user ServiceErrCode
func ConvertErr(err error) ErrNo {
	errno := ErrNo{}
	if errors.As(err, &errno) {
		return errno
	}

	s := ServiceError
	s.ErrorMsg = err.Error()
	return s
}
