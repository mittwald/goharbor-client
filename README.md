# goharbor-client

![Compile & Test](https://github.com/mittwald/goharbor-client/workflows/Compile%20&%20Test/badge.svg)
[![GitHub license](https://img.shields.io/github/license/mittwald/goharbor-client.svg)](https://github.com/mittwald/goharbor-client/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/mittwald/goharbor-client)](https://goreportcard.com/report/github.com/mittwald/goharbor-client)
[![Documentation](https://godoc.org/github.com/mittwald/goharbor-client?status.svg)](https://pkg.go.dev/github.com/mittwald/goharbor-client)

<hr />

:warning: **STABILITY NOTICE:** This project is still under development and not stable yet - breaking changes may happen at any time and without notice

<hr/>

A Harbor API client enabling Go programs to perform CRUD operations on [goharbor](https://github.com/goharbor/harbor) users and projects.

This library is built upon `goharbor/v1.10.2`, and utilizes typings from the upstream source [goharbor/harbor](https://github.com/goharbor/harbor), available under the 
[Apache 2 license](https://github.com/goharbor/harbor/blob/master/LICENSE).

The initial project is a fork of [TimeBye/go-harbor](https://github.com/TimeBye/go-harbor) and available under the MIT License.

## Installation

Install this library using `go get`:

    $ go get github.com/mittwald/goharbor-client

## Usage

Initialize a new goharbor client, then use the various services on the client to
access different parts of the Harbor API.

```go
package main
import (
	"errors"
	"fmt"
	harbor "github.com/mittwald/goharbor-client"
)

func main() {
	client, err := harbor.NewClient("url", "username", "password")
	if err != nil {
		panic(err)
	}

	// Projects
	projects, err := client.Projects().ListProjects(harbor.ListProjectsOptions{})
	if err != nil {
		var e *harbor.StatusCodeError
		if errors.As(err, &e) {
			// handle status code error
			fmt.Printf("request failed with status code: %d", e.StatusCode)
		} else {
			panic(err)
		}
	}

	for _, p := range projects {
		fmt.Println(p.Name)
	}
	
	// Users
	users, err := client.Users().ListUsers()
	// ...

	// Replications
	replications, err := client.Replications().GetReplicationExecutionsByID(1)
	//
	
	// Registries
	registries, err := client.Registries().GetRegistryByID(1)
	//

	...
}
```

## Documentation
For more specific documentation, please refer to the [godoc](https://pkg.go.dev/github.com/mittwald/goharbor-client) of this library.
