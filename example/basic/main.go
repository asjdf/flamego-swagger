package main

import (
	flamegoSwagger "github.com/asjdf/flamego-swagger"
	"github.com/flamego/flamego"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger/example/basic/api"

	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io:8080
// @BasePath /v2
func main() {
	r := flamego.New()

	r.Get("/v2/testapi/get-string-by-int/:some_id", api.GetStringByInt)
	r.Get("/v2/testapi/get-struct-array-by-string/:some_id", api.GetStructArrayByString)

	url := flamegoSwagger.URL("http://petstore.swagger.io:8080/swagger/doc.json") // The url pointing to API definition
	r.Get("/swagger/{**}", flamegoSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run()
}
