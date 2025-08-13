package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/ziflex/lecho/v3"
)

const (
	errInternal   = "500 Internal Server Error"
	errBadRequest = "400 Bad Request"

	newUserScore = 100
)

func main() {
	e := echo.New()
	e.Logger = lecho.New(os.Stdout)

	handler := &echoHandler{}
	mockHandler := &notImplementedHandler{}

	e.POST("/api/themes", handler.createTheme)
	e.GET("/api/themes", mockHandler.notImplemented)
	e.GET("/api/themes/:id", mockHandler.notImplemented)
	e.PUT("/api/themes/:id", mockHandler.notImplemented)
	e.DELETE("/api/themes/:id", mockHandler.notImplemented)

	e.POST("/api/books", handler.createBook)
	e.GET("/api/books", mockHandler.notImplemented)
	e.GET("/api/books/:id", mockHandler.notImplemented)
	e.PUT("/api/books/:id", mockHandler.notImplemented)
	e.DELETE("/api/books/:id", mockHandler.notImplemented)

	go func() {
		if err := e.Start(":12345"); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	gs := <-quit

	signal.Stop(quit)
	log.Info().Str("signal", gs.String()).Msg("Shutting down.")
}
