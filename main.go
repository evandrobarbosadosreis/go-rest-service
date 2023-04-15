package main

import (
	"github.com/evandrobarbosadosreis/go-rest-development/logger"
	"github.com/evandrobarbosadosreis/go-rest-development/web"
)

func main() {
	logger.Info("Starting application...")
	web.Start()
}
