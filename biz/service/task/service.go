package taskService

import (
	"nodolist/biz/dal"
	"nodolist/biz/model/task"
	"nodolist/pkg/errno"
	"time"
)

func (s *TaskService) CreateTask(req *task.CreateTaskReq, uid int64) (dal.Task, error) {
	taskModel := new(dal.Task)
	taskModel = &dal.Task{
		Uid:     uid,
		Title:   req.Title,
		Content: req.Content,
		Status:  0,
	}
	return dal.CreateTask(s.ctx, taskModel)
}

func (s *TaskService) ChangeTaskStatus(req *task.ChangeTaskStatusReq, uid int64) error {
	taskModel := new(dal.Task)
	taskModel = &dal.Task{
		ID:     req.ID,
		Uid:    uid,
		Status: req.Status,
	}
	taskResp, ok, err := dal.QueryTaskByID(s.ctx, taskModel)
	if err != nil {
		return err
	}
	if !ok {
		return errno.TaskNotExistError
	}
	if taskResp.Status == taskModel.Status {
		return errno.TaskStatusError
	}
	switch req.Status {
	case 1:
		taskModel.FinishedAt = time.Now().Format("2006-01-02 15:01:04")
	case 0:
		taskModel.FinishedAt = ""
	}
	return dal.ChangeTaskStatus(s.ctx, taskModel)
}

func (s *TaskService) ShowDoingTasks(req *task.ShowNotDoTaskReq, uid int64) (*[]dal.Task, int64, error) {
	return dal.GetDoingTasksByStatus(int(req.PageNum), uid, 0)
}

func (s *TaskService) ShowDoneTasks(req *task.ShowDoneTaskReq, uid int64) (*[]dal.Task, int64, error) {
	return dal.GetDoingTasksByStatus(int(req.PageNum), uid, 1)
}
