package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func saveUser(c echo.Context) error {
	var req createUserReq
	err := c.Bind(&req)
	if err != nil {
		return c.String(http.StatusInternalServerError, errInternal)
	}

	if strings.TrimSpace(req.ID) == "" {
		return c.String(http.StatusBadRequest, errBadRequest)
	}

	users.mu.Lock()
	defer users.mu.Unlock()

	_, ok := users.db[req.ID]
	if ok {
		return c.String(http.StatusBadRequest, "User already exists.")
	}

	users.db[req.ID] = user{
		ID:    req.ID,
		Score: newUserScore,
	}

	return c.JSON(http.StatusCreated, map[string]any{"newUserId": req.ID})
}

func getUser(c echo.Context) error {
	panic("notReady")
}

func updateUser(c echo.Context) error {
	panic("notReady")
}

func deleteUser(c echo.Context) error {
	panic("notReady")
}
