package main

import (
	"fmt"
	"job-queue/pkg/shared"
	"job-queue/pkg/worker"
)

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

	w := worker.New(config)
	w.Run()

}
