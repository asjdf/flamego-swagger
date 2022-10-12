package v1

import (
	"github.com/flamego/flamego"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1

func Register(router *flamego.Flame) {
	router.Group("/v1", func() {
		router.Get("/books", GetBooks)
	})
}
