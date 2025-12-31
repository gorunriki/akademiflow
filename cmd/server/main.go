package main

import (
	"fmt"

	"github.com/gorunriki/akademiflow/pkg/config"
)

func main() {
	cfg := config.Load()
	fmt.Println("App:", cfg.AppName)
	fmt.Println("Running on port", cfg.AppPort)
}
