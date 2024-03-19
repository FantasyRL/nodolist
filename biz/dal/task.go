package dal

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"nodolist/pkg/conf"
	"time"
)

type Task struct {
	ID         int64
	Uid        int64
	Title      string
	Content    string
	Status     int64
	FinishedAt string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `sql:"index"`
}

func QueryTaskByID(ctx context.Context, task *Task) (*Task, bool, error) {
	taskResp := new(Task)
	err := DB.Model(Task{}).Where("id = ? AND uid = ?", task.ID, task.Uid).First(taskResp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	return taskResp, true, nil
}

func CreateTask(ctx context.Context, task *Task) (Task, error) {
	if err := DB.WithContext(ctx).Create(task).Error; err != nil {
		return Task{}, err
	}
	return *task, nil
}

func ChangeTaskStatus(ctx context.Context, task *Task) error {
	if err := DB.Model(Task{}).Where("id = ?", task.ID).Update("status", task.Status).Update("finished_at", task.FinishedAt).Error; err != nil {
		return err
	}

	return nil
}

func GetDoingTasksByStatus(pageNum int, uid int64, status int64) (*[]Task, int64, error) {
	tasks := new([]Task)
	var count int64
	if err := DB.Model(Task{}).Where("uid = ? AND status = ?", uid, status).Count(&count).Limit(conf.PageSize).Offset((pageNum - 1) * conf.PageSize).Find(tasks).Error; err != nil {
		return nil, 0, err
	}
	return tasks, count, nil
}

func GetTodoCountByUid(uid int64) int64 {
	var count int64
	if err := DB.Model(Task{}).Where("uid = ? AND status = 0", uid).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0
		}
		log.Println(err)
		return 0
	}
	return count
}
