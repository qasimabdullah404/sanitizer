package main

import (
	"fmt"
	"net/http"

	"gihtub.com/qasimabdullah404/sanitizer/internal/config"
	"gihtub.com/qasimabdullah404/sanitizer/internal/handler"
	"gihtub.com/qasimabdullah404/sanitizer/internal/logger"
	"gihtub.com/qasimabdullah404/sanitizer/internal/middleware"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel)
	defer log.Sync()

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware(log))
	r.HandleFunc("/sanitize", handler.UploadHandler(log)).Methods("POST")

	port := fmt.Sprintf(":%s", cfg.Port)
	log.Info("Starting server", zap.String("port", port))
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal("Server failed", zap.Error(err))
	}
}
