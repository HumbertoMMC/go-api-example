package main

import (
	"github.com/akrylysov/algnhsa"
	"github.com/h-m/rest-api/handlers"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Logger was not created")
	}

	h := handlers.NewHandler(logger)
	algnhsa.ListenAndServe(h, nil)
}
