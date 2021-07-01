# Contribution guide

## Table of contents
- [Implementing sub-clients](#implementing-sub-clients)
- [Code generation](#code-generation)
    - [Client APIs](#client-apis)
    - [API mocks](#api-mocks)

## Implementing sub-clients
This client library includes separate clients supporting the `v1.10` & `v2.1` versions of `goharbor`.

This section helps you to get started when adding / updating a new sub-client.

--- 
A new sub-client can be added by creating the corresponding package in either the [apiv1](./apiv1) or [apiv2](./apiv2) directory.

For example, let's look at the existing [`user` sub-client package](https://github.com/mittwald/goharbor-client/tree/master/apiv2/user) of the `v2` client:
```shell
.
└── apiv2
  └── user
      ├── user.go
      ├── user_errors.go
      ├── user_integration_test.go
      └── user_test.go
```
> To maintain integrity with the rest of the repository,
 new sub-client's `.go`-files should be prefixed with their package name.

`user.go` holds the methods that act operations on `user` objects on the `goharbor` API.

It contains a [`RESTClient` struct](https://github.com/mittwald/goharbor-client/blob/master/apiv2/user/user.go#L17) that - for compatibility reasons - groups together the **legacy and v2** client APIs
as well as a field `AuthInfo` for [openAPI's `runtime.ClientAuthInfoWriter`](https://pkg.go.dev/github.com/go-openapi/runtime#ClientAuthInfoWriter).
The latter is used to authenticate requests to the `goharbor` API:
```go
import (
    "context"
    "errors"
    "github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client"
    "github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client/products"
    v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
    model "github.com/mittwald/goharbor-client/v4/apiv2/model/legacy"
    "github.com/go-openapi/runtime"
)

// RESTClient is a subclient for handling user related actions.
type RESTClient struct {
	// The legacy swagger client
	LegacyClient *client.Harbor

	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}
```

A [`NewClient` function](https://github.com/mittwald/goharbor-client/blob/master/apiv2/user/user.go#L28) constructs and returns the `user.RESTClient`:
```go
func NewClient(legacyClient *client.Harbor, v2Client *v2client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		LegacyClient: legacyClient,
		V2Client:     v2Client,
		AuthInfo:     authInfo,
	}
}

```

The package also contains a [`Client` interface](https://github.com/mittwald/goharbor-client/blob/master/apiv2/user/user.go#L36) that holds all method signatures:
```go
type Client interface {
	[...]
}

```

Since there are no methods in this example setup _yet_, let's implement a method called `GetUser` that returns a `user` object:
```go
// GetUser returns an existing user or an error in case of failure.
func (c *RESTClient) GetUser(ctx context.Context, username string) (*model.User, error) {
    if username == "" {
        return nil, errors.New("no username provided")
    }
    
    resp, err := c.LegacyClient.Products.GetUsers(&products.GetUsersParams{
        Context:  ctx,
        Username: &username,
    }, c.AuthInfo)
    if err != nil {
        return nil, handleSwaggerUserErrors(err)
    }
    
    for _, v := range resp.Payload {
        if v.Username == username {
            return v, nil
        }
    }
    
    return nil, &ErrUserNotFound{}
}
```

The `Client` interface should now change to include the method signature of `GetUser`:
```go
type Client interface {
	GetUser(ctx context.Context, username string) (*model.User, error)
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

// GetUser wraps the GetUser method of the user sub-package.
func (c *RESTClient) GetUser(ctx context.Context, username string) (*model.User, error) {
    return c.user.GetUser(ctx, username)
}
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
