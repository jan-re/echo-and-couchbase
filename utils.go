package main

import (
	"math/rand"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	idLen     = 8
	idCharset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

type notImplementedHandler struct{}

func (nih *notImplementedHandler) notImplemented(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "Endpoint not implemented")
}

func generateID() string {
	sb := strings.Builder{}
	sb.Grow(idLen)

	for i := 0; i < idLen; i++ {
		num := rand.Intn(len(idCharset))
		sb.WriteByte(idCharset[num])
	}

	return sb.String()
}
