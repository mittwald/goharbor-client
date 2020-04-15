# go-harbor

A Harbor API client enabling Go programs to perform CRUD operations on [goharbor](https://github.com/goharbor/) users and projects

[![GitHub license](https://img.shields.io/github/license/elenz97/go-harbor.svg)](https://github.com/elenz97/go-harbor/blob/master/LICENSE)

This library is mainly build upon `goharbor/v1.10.1`

## Usage

Initialize a new go-harbor client, then use the various services on the client to
access different parts of the Harbor API.

```go
harborClient := harbor.NewClient(nil, "url","username","password")
```

## Documentation
For more specific documentation, please refer to the [godoc](https://godoc.org/github.com/elenz97/go-harbor) of this library
