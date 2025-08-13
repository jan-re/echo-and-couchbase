package main

import (
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/labstack/echo/v4"
)

// TODO Delete this once no longer useful, should be replaced by DB interface
var users = struct {
	mu sync.Mutex
	db map[string]user
}{
	db: map[string]user{},
}

// TODO Delete this once no longer useful
func saveUser(c echo.Context) error {
	var req createUserReq
	err := c.Bind(&req)
	if err != nil {
		return c.String(http.StatusInternalServerError, errInternal)
	}

	if strings.TrimSpace(req.ID) == "" {
		return c.String(http.StatusBadRequest, errBadRequest)
	}

	_, ok := users.db[req.ID]
	if ok {
		return c.String(http.StatusBadRequest, "User already exists.")
	}

	users.mu.Lock()
	defer users.mu.Unlock()

	users.db[req.ID] = user{
		ID:    req.ID,
		Score: newUserScore,
	}

	return c.JSON(http.StatusCreated, map[string]any{"userId": req.ID})
}

// TODO Delete this once no longer useful
func getUser(c echo.Context) error {
	var req getUserReq
	err := c.Bind(&req)
	if err != nil {
		return c.String(http.StatusInternalServerError, errInternal)
	}

	user, ok := users.db[req.ID]
	if !ok {
		return c.String(http.StatusNotFound, "User not found.")
	}

	resp := map[string]any{"userId": user.ID}

	var getScore bool
	if strings.TrimSpace(req.GetScore) != "" {
		getScore, err = strconv.ParseBool(req.GetScore)
		if err != nil {
			return c.String(
				http.StatusBadRequest,
				"Query paramater \"getScore\" value must be true or false. Other values are not permitted.",
			)
		}
	}

	if getScore {
		resp["score"] = user.Score
	}

	return c.JSON(http.StatusOK, resp)
}

// TODO Delete this once no longer useful
func updateUser(c echo.Context) error {
	var req putUserReq
	err := c.Bind(&req)
	if err != nil {
		return c.String(http.StatusInternalServerError, errInternal)
	}

	user, ok := users.db[req.ID]
	if !ok {
		return c.String(http.StatusNotFound, "User not found.")
	}

	var updated bool
	if req.Name != nil {
		user.Name = *req.Name
		updated = true
	}

	if req.Score != nil {
		user.Score = *req.Score
		updated = true
	}

	users.mu.Lock()
	defer users.mu.Unlock()

	users.db[req.ID] = user

	return c.JSON(http.StatusOK, map[string]any{"userUpdated": updated})
}
