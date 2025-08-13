package main

import (
	"net/http"
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

	e.POST("/api/themes", notImplemented)
	e.GET("/api/themes", notImplemented)
	e.GET("/api/themes/:id", notImplemented)
	e.PUT("/api/themes/:id", notImplemented)
	e.DELETE("/api/themes/:id", notImplemented)

	e.POST("/api/books", notImplemented)
	e.GET("/api/books", notImplemented)
	e.GET("/api/books/:id", notImplemented)
	e.PUT("/api/books/:id", notImplemented)
	e.DELETE("/api/books/:id", notImplemented)

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

func notImplemented(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "Endpoint not implemented")
}
