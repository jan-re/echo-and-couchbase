package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

const (
	errInternal   = "500 Internal Server Error"
	errBadRequest = "400 Bad Request"

	newUserScore = 100
)

var users = struct {
	mu sync.Mutex
	db map[string]user
}{
	db: map[string]user{},
}

func main() {
	e := echo.New()

	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	go func() {
		if err := e.Start(":12345"); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	gs := <-quit

	signal.Stop(quit)
	log.Info().Str("signal", gs.String()).Msg("Terminating.")
}
