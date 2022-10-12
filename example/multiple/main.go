package main

import (
	flamegoSwagger "github.com/asjdf/flamego-swagger"
	v1 "github.com/asjdf/flamego-swagger/example/multiple/api/v1"
	v2 "github.com/asjdf/flamego-swagger/example/multiple/api/v2"
	_ "github.com/asjdf/flamego-swagger/example/multiple/docs"
	"github.com/flamego/flamego"
	swaggerFiles "github.com/swaggo/files"
)

func main() {
	// New flamego router
	router := flamego.New()

	// Register api/v1 endpoints
	v1.Register(router)
	router.Get("/swagger/v1/*any", flamegoSwagger.WrapHandler(swaggerFiles.NewHandler(), flamegoSwagger.InstanceName("v1")))

	// Register api/v2 endpoints
	v2.Register(router)
	router.Get("/swagger/v2/*any", flamegoSwagger.WrapHandler(swaggerFiles.NewHandler(), flamegoSwagger.InstanceName("v2")))

	// Listen and Server in
	router.Run()
}
