namespace go task

include"user.thrift"
include"common.thrift"

struct Task{
    1:user.User user,
    2:i64 id,
    3:string title,
    4:string content,
    5:i64 status,
    6:string created_at,
    7:string finished_at,
}

struct CreateTaskReq{
    1:string title,
    2:string content,
}

struct CreateTaskResp{
    1:common.BaseResp base,
    2:Task task,
}