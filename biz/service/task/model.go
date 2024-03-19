package taskService

import (
	"context"
	"nodolist/biz/dal"
	"nodolist/biz/model/task"
)

type TaskService struct {
	ctx context.Context
}

func NewTaskService(ctx context.Context) *TaskService {
	return &TaskService{ctx: ctx}
}

func BuildTaskResp(taskDao dal.Task) *task.Task {
	//userResp, err := dal.QueryUserByID(taskDao.Uid)
	//if err != nil {
	//	log.Println(err)
	//}
	return &task.Task{
		//User:       userService.BuildUserResp(userResp),
		ID:         taskDao.ID,
		Title:      taskDao.Title,
		Content:    taskDao.Content,
		Status:     taskDao.Status,
		CreatedAt:  taskDao.CreatedAt.Format("2006-01-02 15:01:04"),
		FinishedAt: taskDao.FinishedAt,
	}
}

func BuildTasksResp(tasksDao *[]dal.Task) []*task.Task {
	var taskListResp []*task.Task
	for _, taskResp := range *tasksDao {
		taskListResp = append(taskListResp, &task.Task{
			ID:         taskResp.ID,
			Title:      taskResp.Title,
			Content:    taskResp.Content,
			CreatedAt:  taskResp.CreatedAt.Format("2006-01-02 15:01:04"),
			FinishedAt: taskResp.FinishedAt,
		})
	}
	return taskListResp
}
