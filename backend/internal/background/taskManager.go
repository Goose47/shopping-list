package background

import (
	"gorm.io/gorm"
	"log/slog"
)

type TaskManager struct {
	log *slog.Logger
	db  *gorm.DB
}

func NewTaskManager(log *slog.Logger, db *gorm.DB) *TaskManager {
	return &TaskManager{
		log: log,
		db:  db,
	}
}

func (tm *TaskManager) Run() {
}
