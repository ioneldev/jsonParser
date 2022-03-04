package main

import (
	"fmt"
	"jsonParser/processor"
	"jsonParser/service"
)

func main() {
	fmt.Println("Started")
	config := service.Config{}

	err := config.Initialize("config.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	worker := processor.Worker{}
	worker.Initialize(config)
	err = worker.Start()

	if err != nil {
		fmt.Println(err)
		return
	}
}
