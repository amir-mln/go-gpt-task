package main

import (
	"fmt"
	"go-gpt-task/configs"
)

func main() {
	cfg, err := configs.Load()
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg)
}
