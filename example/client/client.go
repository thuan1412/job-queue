package main

import (
	"context"
	"fmt"
	"job-queue/pkg/client"
	"job-queue/pkg/shared"
	"math/rand"
	"time"
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

	for {
		_, err := c.AddTask(context.Background(), "test", TaskArgs{rand.Int()})
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}

}
