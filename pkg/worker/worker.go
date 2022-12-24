package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"job-queue/pkg/constants"
	"job-queue/pkg/shared"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

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
		rows, err := w.db.QueryContext(context.Background(), "SELECT * FROM task_queue WHERE status = ?", constants.STATUS_PENDING)
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			task := Task{}
			err := rows.Scan(&task.ArgsStr, &task.UUID, &task.Status, &task.TaskType, &task.ID)
			panicE(err)

			json.Unmarshal([]byte(task.ArgsStr), &task.Args)
			fmt.Println(task)

			if handler, ok := w.config.HandlerMap[task.TaskType]; ok {
				err := handler(task.Args)
				if err != nil {

				}
			}
		}
		time.Sleep(time.Second)
	}
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
