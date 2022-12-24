package client

import (
	"context"
	"encoding/json"
	"job-queue/pkg/constants"
	"job-queue/pkg/shared"

	"database/sql"

	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

type Client struct {
	config shared.Config
	db     *sql.DB
}

func (c Client) AddTask(ctx context.Context, taskType string, args any) (string, error) {
	uid := uuid.New().String()
	argsStr, err := json.Marshal(args)
	if err != nil {
		return "", err
	}
	_, err = c.db.ExecContext(ctx, "INSERT INTO task_queue (uuid, args, task_type, status) VALUES (?, ?, ?, ?);", uid, argsStr, taskType, constants.STATUS_PENDING)
	return uid, err
}

func New(config shared.Config) Client {
	db, err := sql.Open("mysql", config.Url)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return Client{
		config,
		db,
	}
}
