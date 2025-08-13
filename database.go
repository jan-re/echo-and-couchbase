package main

import (
	"context"
	"errors"
)

var (
	errAlreadyExists = errors.New("document already exists")
)

type database interface {
	storeTheme(context.Context, theme) error
	storeBook(context.Context, book) error

	// TODO Fill in methods to suit the handled routes
}
