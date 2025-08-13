package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type echoHandler struct {
	database database
}

func (eh *echoHandler) createTheme(c echo.Context) error {
	var newTheme theme
	err := c.Bind(&newTheme)
	if err != nil {
		// TODO Improve messaging, notify the caller about err reason
		return echo.ErrBadRequest
	}

	if strings.TrimSpace(newTheme.Name) == "" {
		// TODO Improve messaging, notify the caller about err reason
		return echo.ErrBadRequest
	}

	newTheme.ID = generateID()

	// TODO Assure that [theme.Name] uniqueness is enforced by a database constrant.
	err = eh.database.storeTheme(c.Request().Context(), newTheme)
	if err != nil {
		if errors.Is(err, errAlreadyExists) {
			// TODO Improve messaging, notify the caller about err reason
			return echo.ErrBadRequest
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, newTheme)
}

func (eh *echoHandler) createBook(c echo.Context) error {
	var newBook book
	err := c.Bind(&newBook)
	if err != nil {
		// TODO Improve messaging, notify the caller about err reason
		return echo.ErrBadRequest
	}

	if strings.TrimSpace(newBook.Name) == "" {
		// TODO Improve messaging, notify the caller about err reason
		return echo.ErrBadRequest
	}

	if strings.TrimSpace(newBook.Author) == "" {
		// TODO Improve messaging, notify the caller about err reason
		return echo.ErrBadRequest
	}

	if newBook.PageCount <= 0 {
		// TODO Improve messaging, notify the caller about err reason
		return echo.ErrBadRequest
	}

	newBook.ID = generateID()

	// TODO Assure that [book.Name] for the given [book.Author] uniqueness is enforced by a database constrant.
	err = eh.database.storeBook(c.Request().Context(), newBook)
	if err != nil {
		if errors.Is(err, errAlreadyExists) {
			// TODO Improve messaging, notify the caller about err reason
			return echo.ErrBadRequest
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, newBook)
}
