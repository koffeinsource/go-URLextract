package main

import (
	"net/http"

	"github.com/koffeinsource/go-klogger"
)

// Config is the struct used to configure go-URLextract
type config struct {
	Client *http.Client
	Logger *klogger.KLogger
}

func (c *config) HTTPClient() *http.Client {
	return c.Client
}
func (c *config) Log() *klogger.KLogger {
	return c.Logger
}
