package main

import (
	"context"
	"fmt"
	"job-queue/pkg/client"
	"job-queue/pkg/shared"
	"job-queue/pkg/worker"
)

type TaskArgs struct {
	Id int `json:"id"`
}

func main() {
	dbUrl := "dev:dev@tcp(127.0.0.1:3306)/test_employee"
	config := shared.Config{
		Url: dbUrl,
		HandlerMap: map[string]shared.Handler{
			"test": func(args any) error {
				fmt.Println(args)
				return nil
			},
		},
	}

	c := client.New(config)

	_, err := c.AddTask(context.Background(), "test", TaskArgs{1})

	w := worker.New(config)
	w.Run()

	if err != nil {
		panic(err)
	}
}
