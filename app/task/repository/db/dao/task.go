package dao

import (
	"context"
	"rpcoptool/app/task/repository/db/model"
	"rpcoptool/proto/pb"

	"gorm.io/gorm"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	return &TaskDao{NewDBClient(ctx)}
}

func (dao *TaskDao) ListTask(start, limit int) (r []*model.Task, count int64, err error) {
	err = dao.Model(&model.Task{}).Offset(start).
		Limit(limit).
		Find(&r).Error
	if err != nil {
		return
	}
	err = dao.Model(&model.Task{}).
		Count(&count).Error

	return
}
func (dao *TaskDao) CreateTask(in *model.Task) error {
	return dao.Model(&model.Task{}).Create(&in).Error

}
func (dao *TaskDao) CreateOpT(in *model.TaskOp) error {
	return dao.Model(&model.TaskOp{}).Create(&in).Error
}
func (dao *TaskDao) GetTaskById(taskId uint64) (r *model.Task, err error) {
	err = dao.Model(&model.Task{}).
		Where("id = ?", taskId).
		First(&r).Error
	return
}
func (dao *TaskDao) UpdateTask(in *pb.TaskModel) (err error) {
	task := new(model.Task)
	err = dao.Model(&model.Task{}).
		Where("id = ?", in.Id).
		First(&task).Error
	if err != nil {
		return
	}
	task.Content = in.Content
	err = dao.Save(&task).Error
	return
}
