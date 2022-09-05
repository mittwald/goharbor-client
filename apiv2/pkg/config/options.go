package config

import (
	"time"
)

// Options defines optional parameters for configuring an API client.
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
		Page:     1,
		Sort:     "",
		Query:    "",
		Timeout:  30 * time.Second,
	}
}

func (o *Options) WithPageSize(pageSize int64) *Options {
	o.PageSize = pageSize
	return o
}

func (o *Options) WithPage(page int64) *Options {
	o.Page = page
	return o
}

func (o *Options) WithTimeout(timeout time.Duration) *Options {
	o.Timeout = timeout
	return o
}

func (o *Options) WithSort(sort string) *Options {
	o.Sort = sort
	return o
}

func (o *Options) WithQuery(query string) *Options {
	o.Query = query
	return o
}
