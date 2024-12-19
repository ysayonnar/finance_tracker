package main

import (
	"financeTracker/pkg/config"
	"fmt"
)

func main() {
	cfg, err := config.ParseConfig()
	if err != nil {
		fmt.Println(err)
	}

	_ = cfg
}
