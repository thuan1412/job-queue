package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"job-queue/pkg/constants"
	"job-queue/pkg/shared"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func warnE(err error) {
	if err != nil {
		log.Println("[WARN]", err)
	}
}

func panicE(err error) {
	if err != nil {
		panic(err)
	}
}

type Worker struct {
	config shared.Config
	db     *sql.DB
}

func (w *Worker) Run() {
	for {
		rows, err := w.db.QueryContext(context.Background(), "SELECT * FROM task_queue WHERE status = ? LIMIT 10;", constants.STATUS_PENDING)
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			task := Task{}
			err := rows.Scan(&task.ArgsStr, &task.UUID, &task.Status, &task.TaskType, &task.ID)
			panicE(err)

			json.Unmarshal([]byte(task.ArgsStr), &task.Args)

			if handler, ok := w.config.HandlerMap[task.TaskType]; ok {
				err := handler(task.Args)
				if err != nil {
					err := w.UpdateTaskStatus(task.UUID, constants.STATUS_PENDING)
					warnE(err)
				}
				err = w.UpdateTaskStatus(task.UUID, constants.STATUS_SUCCESS)
				warnE(err)
			}
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func (w *Worker) UpdateTaskStatus(uuid string, status string) error {
	_, err := w.db.Exec("UPDATE task_queue SET status = ? WHERE uuid = ?", status, uuid)
	return err
}

func New(config shared.Config) *Worker {
	db, err := sql.Open("mysql", config.Url)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return &Worker{
		config,
		db,
	}
}
