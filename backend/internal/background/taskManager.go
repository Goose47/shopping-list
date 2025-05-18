package background

import (
	"fmt"
	"gorm.io/gorm"
	"log/slog"
	"time"
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
	go tm.test()
}

func (tm *TaskManager) test() {
	for {
		fmt.Println("test..")
		time.Sleep(5 * time.Second)
	}
}
