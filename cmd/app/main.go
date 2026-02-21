package main

import (
	"nosql-course/cmd/internal/config"
	"nosql-course/cmd/internal/handler/http"
	"nosql-course/cmd/internal/logger"

	"github.com/labstack/echo/v5"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	log, err := logger.New(cfg.Logger)
	if err != nil {
		panic(err)
	}

	h := http.NewHandler()

	e := echo.New()
	e.Logger = log
	e.GET("/health", h.Health)

	if err = e.Start(cfg.Addr()); err != nil {
		log.Error("shutdown server with error", "error", err)
	}
}
