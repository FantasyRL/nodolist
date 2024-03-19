namespace go user

include "common.thrift"

struct User{
    1: i64 id,
    2: string name,
    3: i64 todo_count,
}

struct RegisterReq {
    1: string username,
//    2: string email,
    2: string password,
}

struct RegisterResp {
    1: common.BaseResp base,
    2: i64 user_id,
}

struct LoginReq {
    1: string username,
    2: string password,
//    3: optional string otp,
}

struct LoginResp {
    1: common.BaseResp base,
    2: User user,
    3: string token,
}

struct InfoReq {
}

struct InfoResp {
    1: common.BaseResp base,
    2: User user,
}

service UserHandler {
    RegisterResp Register(1: RegisterReq req)(api.post="/nodolist/user/register/"),
    LoginResp Login(1: LoginReq req)(api.post="/nodolist/user/login/"),
    InfoResp Info(1: InfoReq req)(api.get="/nodolist/user/"),
//    OTP2FAResp OTP2FA(1:OTP2FAReq req)(api.get="/bibi/user/2fa"),
//    Switch2FAResp Switch2FA(1:Switch2FAReq req)(api.get="/bibi/user/switch2fa"),
}