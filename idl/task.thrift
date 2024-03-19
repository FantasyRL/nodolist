namespace go task

include"user.thrift"
include"common.thrift"

struct Task{
    1:optional user.User user,
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
    2:optional Task task,
}

struct ChangeTaskStatusReq{
    1:i64 id,
    2:i64 status,
}

struct ChangeTaskStatusResp{
    1:common.BaseResp base,
}

struct ShowNotDoTaskReq{
    1:i64 page_num,
}

struct ShowNotDoTaskResp{
    1:common.BaseResp base,
    2:optional i64 task_count,
    3:optional list<Task> tasks,
}

struct ShowDoneTaskReq{
    1:i64 page_num,
}

struct ShowDoneTaskResp{
    1:common.BaseResp base,
    2:optional i64 task_count,
    3:optional list<Task> tasks,
}

service TaskHandler{
    CreateTaskResp CreateTask(1:CreateTaskReq req)(api.post="/nodolist/task/create"),
    ChangeTaskStatusResp ChangeTaskStatus(1:ChangeTaskStatusReq req)(api.get="/nodolist/task/status"),
    ShowNotDoTaskResp ShowNotDoTask(1:ShowNotDoTaskReq req)(api.get="/nodolist/tasks/doing"),
    ShowDoneTaskResp ShowDoneTask(1:ShowDoneTaskReq req)(api.get="/nodolist/tasks/done"),
}