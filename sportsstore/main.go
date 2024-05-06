package main

import (
	"platform/http"
	"platform/http/handling"
	"platform/pipeline"
	"platform/pipeline/basic"
	"platform/services"
	"sportsstore/models/repo"
	"sportsstore/store"
	"sync"
)

func registerServices() {
	services.RegisterDefaultServices()
	repo.RegisterMemoryRepoService()
}

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServiceComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.StaticFileComponent{},
		handling.NewRouter(
			handling.HandlerEntry{Prefix: "", Handler: store.ProductHandler{}},
		).AddMethodsAlias("/", store.ProductHandler.GetProducts, 1).
			AddMethodsAlias("/products", store.ProductHandler.GetProducts, 1),
	)
}
func main() {
	registerServices()
	results, err := services.Call(http.Serve, createPipeline())
	if err == nil {
		(results[0].(*sync.WaitGroup)).Wait()
	} else {
		panic(err)
	}
}
