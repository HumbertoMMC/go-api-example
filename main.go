package main

import (
	"net/http"

	"github.com/h-m/rest-api/handlers"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Logger was not created")
	}
	http.Handle("/", handlers.NewHandler(logger))
	logger.Info("Server listening on port 8000")
	http.ListenAndServe("localhost:8000", nil)
}
