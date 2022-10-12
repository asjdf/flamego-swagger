package v2

import (
	"github.com/flamego/flamego"
)

// @title Swagger Example API
// @version 2.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v2

func Register(router *flamego.Flame) {
	router.Group("/v2", func() {
		router.Get("/books", GetBooks)
	})
}
