# Contribution guide

## Table of contents
- [Repository structure](#repository-structure)
- [Implementing sub-clients](#implementing-sub-clients)
- [Code generation](#code-generation)
    - [Client APIs](#client-apis)
    - [API mocks](#api-mocks)

## Repository structure

```shell
.
├── apiv1
│  ├── client.go //  The main file of the package containing the exposed 'Client' interface
│  ├── internal
│  │  └── api
│  │      └── client
│  │          ├── harbor_client.go
│  │          └── products // Contains the 'Products' client which wraps all v1 client functions
│  │              └── products_client.go // contains the ClientService interface used to generate 'client_service.go'
│  ├── mocks
│  │  └── client_service.go
│  ├── model // v1 API definitions
│  ├── project // Project client (as an example)
│  │  ├── project_errors.go
│  │  ├── project.go
│  │  ├── project_integration_test.go
│  │  └── project_test.go
│  └── [...] // Other clients
│  ├── scripts -> ../scripts // Symbolic link to code generation scripts
│  ├── testdata -> ../testdata // Symbolic link to integration test configuration
│  └── testing // Testing definitions
│
├── apiv2
│  ├── client.go // The main file of the package containing the exposed 'Client' interface
│  ├── internal
│  │  └── api
│  │     └── client
│  │         ├── harbor_client.go // Contains the 'Harbor' struct which wraps all of the below clients
│  │         ├── auditlog // 'Auditlog' 'v2' API client wrapping all auditlog functions (as an example)
│  │         └── [...] // Other 'v2' API client functions
│  ├── mocks // Contains mock clients of the above clients in './apiv2/internal/api/client'
│  │  └── [...]
│  ├── model // v2 API definitions
│  │  └── [...]
│  ├── pkg
│  │  ├── clients // Contains the API clients implemented by the 'Client' interface in './apiv2/client.go'
│  │  │  ├── auditlog // 'Auditlog' client (as an example)
│  │  │  │  ├── auditlog_errors.go
│  │  │  │  ├── auditlog.go
│  │  │  │  ├── auditlog_integration_test.go
│  │  │  │  └── auditlog_test.go
│  │  │  └── [...] // Other clients
│  │  ├── common // Types used throughout this package
│  │  │  └── [...]
│  │  ├── config // goharbor-client configuration, e.g. for specifying the 'PageSize' when operating on API endpoints using pagination
│  │  │  └── options.go
│  │  ├── errors // Wrapped errors used throughout this package
│  │  │  └── [...]
│  │  ├── testing // Testing definitions
│  │  └── util
│  │      └── project.go
│  ├── scripts -> ../scripts // Symbolic link to code generation scripts
│  └── testdata -> ../testdata // Symbolic link to integration test configuration
│
├── go.mod
├── go.sum
├── Makefile
├── scripts
│  ├── gen-mock.sh // 'mockery' code generation (used by 'make generate')
│  ├── setup-harbor.sh // Local harbor bootstrapping (requires helm, used by 'make setup-harbor-v1', 'make setup-harbor-v2')
│  └── swagger-gen.sh // 'go-swagger' code generation (used by 'make generate')
└── testdata
    └── kind-config.yaml
```

## Implementing sub-clients

This section helps you to get started when adding / updating a new sub-client.

--- 
A new sub-client can be added by creating the corresponding package in either the [apiv1](./apiv1) or [apiv2/pkg/clients](./apiv2/pkg/clients) directory.

For example, let's look at the existing `user` package of the `v2` client:

```shell
.
└── apiv2
   └── pkg
      └── clients
         └── user
             ├── user_errors.go
             ├── user.go
             ├── user_integration_test.go
             └── user_test.go
```

> To maintain integrity with the rest of the repository,
 new sub-client's `.go`-files should be prefixed with their package name.

`user.go` holds the methods that act operations on `user` objects on the `goharbor` API. The below examples are not guaranteed to be up to date.

It contains a `RESTClient` struct that groups together the `v2` Goharbor client, client `Options` as well as a field `AuthInfo` for [openAPI's `runtime.ClientAuthInfoWriter`](https://pkg.go.dev/github.com/go-openapi/runtime#ClientAuthInfoWriter).
The latter is used to authenticate requests to the `goharbor` API:

```go
package user

import (
	"context"
	"errors"

	"k8s.io/apimachinery/pkg/util/intstr"

	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/user"
	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
	clienterrors "github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"

	"github.com/go-openapi/runtime"
)

// RESTClient is a subclient for handling user related actions.
type RESTClient struct {
	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// Options contains optional configuration when making API calls.
	Options *config.Options

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

```

A `NewClient` function constructs and returns the `user.RESTClient`:

```go
func NewClient(v2Client *v2client.Harbor, opts *config.Options, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		Options:  opts,
		V2Client: v2Client,
		AuthInfo: authInfo,
	}
}
```

The package also contains a [`Client` interface](https://github.com/mittwald/goharbor-client/blob/master/apiv2/user/user.go#L36) that holds all method signatures:
```go
type Client interface {
	[...]
}

```

Since there are no methods in this example setup _yet_, let's implement a method called `GetUserByName` that returns a `*modelv2.UserResp` object:
```go
// GetUserByName returns an existing user identified by name.
func (c *RESTClient) GetUserByName(ctx context.Context, username string) (*modelv2.UserResp, error) {
	if username == "" {
		return nil, errors.New("no username provided")
	}

	c.Options.PageSize = 100

	resp, err := c.ListUsers(ctx)
	if err != nil {
		return nil, handleSwaggerUserErrors(err)
	}

	for _, u := range resp {
		if u.Username == username {
			return u, nil
		}
	}

	return nil, &clienterrors.ErrUserNotFound{}
}

```

The `Client` interface should now change to include the method signature of `GetUser`:
```go
type Client interface {
	GetUserByName(ctx context.Context, username string) (*modelv2.UserResp, error)
	[...]
}

```

---

To make the sub-client's methods accessible via the `github.com/mittwald/goharbor-client/apiv2` package, 
the above implementation has to be wrapped together in [apiv2/client.go](./apiv2/client.go): 

```go
// RESTClient implements the Client interface as a REST client
type RESTClient struct {
	user        *user.RESTClient
    [...]
}

// NewRESTClient constructs a new REST client containing each sub client.
func NewRESTClient(legacyClient *client.Harbor, v2Client *v2client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		user:        user.NewClient(legacyClient, v2Client, authInfo),
        [...]
	}
}

type Client interface {
    user.Client
    [...]
}

// User Client

// GetUserByName wraps the GetUserByName method of the user sub-package.
func (c *RESTClient) GetUserByName(ctx context.Context, username string) (*modelv2.UserResp, error) {
    return c.user.GetUserByName(ctx, username string)
}

[...]
```

## Code generation

### Client APIs
`goharbor-client`'s API types are generated using the `goharbor` [swagger API specifications](https://github.com/goharbor/harbor/tree/master/api).

The API versions used for generation are specified in the repository's [Makefile](./Makefile):
```shell
[...]
V1_VERSION = v1.x.x
V2_VERSION = v2.x.x
```

> Please note, that versions have to match a valid `goharbor` release version. For reference, see: [harbor/releases](https://github.com/goharbor/harbor/releases).

When changing the above API versions, make sure to _always_ run code generation:

```shell
make generate
```

The above versions are then passed on to generator scripts that will fetch the corresponding `.yaml` API specification.

[go-swagger](https://github.com/go-swagger/go-swagger) is then utilized to generate the client API types found under 
`./apiv*/internal`.

### API mocks

`goharbor-client`'s unit tests rely on a mocked goharbor API to interact with.
The [stretchr/testify](https://github.com/stretchr/testify) `assert` & `mock` packages are used throughout all unit tests.

API Mocks for the above types are automatically generated by `make generate`.

After generating new API types, [vektra/mockery](https://github.com/vektra/mockery) is used to generate mocks based on the former.

The resulting mock files can be found under `./apiv*/mocks`.

---

See also:
- [go-swagger/go-swagger](https://github.com/go-swagger/go-swagger)
- [vektra/mockery](https://github.com/vektra/mockery)
- [swagger.io editor](https://editor.swagger.io) (useful for an interactive preview of swagger API's)
- [stretchr/testify](https://github.com/stretchr/testify)
