package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Id        uint   `gorm:"not null"`
	Title     string `gorm:"index; not null"`
	Status    int    `gorm:"default:0"`
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64  `gorm:"default:0"`
	Url       string `gorm:"type:text"`
}
type TaskOp struct {
	gorm.Model
	Id            uint   `gorm:"not null"`
	TaskId        uint   `gorm:"not null"`
	OpId          uint   `gorm:"not null"`
	RpcQueue      string `gorm:"index; not null"`
	Status        int    `gorm:"default:0"`
	CorrelationID string `gorm:"type:longtext"`
}
