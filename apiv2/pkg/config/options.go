package config

import (
	"time"
)

type Options struct {
	// PageSize used for the client operations, a maximum of 100 is enforced by the Goharbor API.
	PageSize int64
	// Page to be used for client operations.
	Page int64
	// The timeout for client operations.
	Timeout time.Duration
	// Sort string used on 'list' client operations.
	Sort string
	// Query string used for client operations.
	Query string
}

func Defaults() *Options {
	return &Options{
		PageSize: 10,
		Page:     0,
		Sort:     "",
		Query:    "",
		Timeout:  30 * time.Second,
	}
}
