# go-harbor

A Harbor API client enabling Go programs to perform CRUD operations on [goharbor](https://github.com/goharbor/harbor) users and projects

[![GitHub license](https://img.shields.io/github/license/elenz97/go-harbor.svg)](https://github.com/elenz97/go-harbor/blob/master/LICENSE)

This library is mainly build upon `goharbor/v1.10.1`

## Usage

Initialize a new go-harbor client, then use the various services on the client to
access different parts of the Harbor API.

```go
package main

import (
    "errors"
    "fmt"
    "github.com/elenz97/go-harbor"
)

func main() {
    client, err := harbor.NewClient("url", "username", "password")
    if err != nil {
        panic(err)
    }

    // Projects
    projects, err := client.Projects().ListProjects(harbor.ListProjectOptions{})
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
}
```

## Documentation
For more specific documentation, please refer to the [godoc](https://pkg.go.dev/github.com/elenz97/go-harbor) of this library
