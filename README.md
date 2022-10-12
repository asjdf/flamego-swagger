# flamego-swagger

flamego middleware to automatically generate RESTful API documentation with Swagger 2.0.

[![Build Status](https://github.com/asjdf/flamego-swagger/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/features/actions)
[![Codecov branch](https://img.shields.io/codecov/c/github/asjdf/flamego-swagger/master.svg)](https://codecov.io/gh/asjdf/flamego-swagger)
[![Go Report Card](https://goreportcard.com/badge/github.com/asjdf/flamego-swagger)](https://goreportcard.com/report/github.com/asjdf/flamego-swagger)
[![GoDoc](https://godoc.org/github.com/asjdf/flamego-swagger?status.svg)](https://godoc.org/github.com/asjdf/flamego-swagger)
[![Release](https://img.shields.io/github/release/asjdf/flamego-swagger.svg?style=flat-square)](https://github.com/asjdf/flamego-swagger/releases)

## Usage

### Start using it

1. Add comments to your API source code, [See Declarative Comments Format](https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format).
2. Download [Swag](https://github.com/swaggo/swag) for Go by using:

```sh
go get -u github.com/swaggo/swag/cmd/swag
```

Starting in Go 1.17, installing executables with `go get` is deprecated. `go install` may be used instead:

```sh
go install github.com/swaggo/swag/cmd/swag
```

3. Run the [Swag](https://github.com/swaggo/swag) at your Go project root path(for instance `~/root/go-peoject-name`),
   [Swag](https://github.com/swaggo/swag) will parse comments and generate required files(`docs` folder and `docs/doc.go`)
   at `~/root/go-peoject-name/docs`.

```sh
swag init
```

4. Download [flamego-swagger](https://github.com/asjdf/flamego-swagger) by using:

```sh
go get -u github.com/asjdf/flamego-swagger
go get -u github.com/swaggo/files
```

Import following in your code:

```go
import "github.com/asjdf/flamego-swagger" // flamego-swagger middleware
import "github.com/swaggo/files" // swagger embed files

```

### Canonical example:

Now assume you have implemented a simple api as following:

```go
// A get function which returns a hello world string by json
func Helloworld(r flamego.Render)  {
	r.JSON(http.StatusOK,"helloworld")
}
```

So how to use flamego-swagger on api above? Just follow the following guide.

1. Add Comments for apis and main function with flamego-swagger rules like following:

```go
// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(r flamego.Render)  {
    r.JSON(http.StatusOK,"helloworld")
}
```

2. Use `swag init` command to generate a docs, docs generated will be stored at `docs/`.
3. import the docs like this:
   I assume your project named `github.com/go-project-name/docs`.

```go
import (
   docs "github.com/go-project-name/docs"
)
```

4. build your application and after that, go to http://localhost:8080/swagger/index.html ,you to see your Swagger UI.

5. The full code and folder relatives here:

```go
package main

import (
   "github.com/flamego/flamego"
   docs "github.com/go-project-name/docs"
   swaggerfiles "github.com/swaggo/files"
   flamegoSwagger "github.com/asjdf/flamego-swagger"
   "net/http"
)
// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(r flamego.Render)  {
   r.JSON(http.StatusOK,"helloworld")
}

func main()  {
   r := flamego.New()
   docs.SwaggerInfo.BasePath = "/api/v1"
   r.Group("/api/v1", func() {
      r.Group("/example", func() {
         r.Get("/helloworld",Helloworld)
      })
   })
   r.Get("/swagger/{**}", flamegoSwagger.WrapHandler(swaggerfiles.Handler))
   r.Run(":8080")

}
```

Demo project tree, `swag init` is run at relative `.`

```
.
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
└── main.go
```

## Multiple APIs
This feature was introduced in swag v1.7.9

## Configuration

You can configure Swagger using different configuration options

```go
func main() {
	r := flamego.New()

	flamegoSwagger.WrapHandler(swaggerFiles.Handler,
		flamegoSwagger.URL("http://localhost:8080/swagger/doc.json"),
		flamegoSwagger.DefaultModelsExpandDepth(-1))

	r.Run()
}
```

| Option                   | Type   | Default    | Description                                                                                                                                                                                                                                                   |
| ------------------------ | ------ | ---------- |---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| URL                      | string | "doc.json" | URL pointing to API definition                                                                                                                                                                                                                                |
| DocExpansion             | string | "list"     | Controls the default expansion setting for the operations and tags. It can be 'list' (expands only the tags), 'full' (expands the tags and operations) or 'none' (expands nothing).                                                                           |
| DeepLinking              | bool   | true       | If set to true, enables deep linking for tags and operations. See the Deep Linking documentation for more information.                                                                                                                                        |
| DefaultModelsExpandDepth | int    | 1          | Default expansion depth for models (set to -1 completely hide the models).                                                                                                                                                                                    |
| InstanceName             | string | "swagger"  | The instance name of the swagger document. If multiple different swagger instances should be deployed on one flamego router, ensure that each instance has a unique name (use the _--instanceName_ parameter to generate swagger documents with _swag init_). |
| PersistAuthorization     | bool   | false      | If set to true, it persists authorization data and it would not be lost on browser close/refresh.                                                                                                                                                             |                                                                                            
| Oauth2DefaultClientID    | string | ""         | If set, it's used to prepopulate the *client_id* field of the OAuth2 Authorization dialog.                                                                                                                                                                    |
