package main

import "errors"

var (
	errAlreadyExists = errors.New("document already exists")
)

type database interface {
	storeTheme(theme) error
	storeBook(book) error

	// TODO Fill in methods to suit the handled routes
}
