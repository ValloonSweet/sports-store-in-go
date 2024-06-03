package main

import (
	"platform/http"
	"platform/http/handling"
	"platform/pipeline"
	"platform/pipeline/basic"
	"platform/services"
	"platform/sessions"
	"sportsstore/admin"
	"sportsstore/models/repo"
	"sportsstore/store"
	"sportsstore/store/cart"
	"sync"
)

func registerServices() {
	services.RegisterDefaultServices()
	// repo.RegisterMemoryRepoService()
	repo.RegisterSqlRepositoryService()
	sessions.RegisterSessionService()
	cart.RegisterCartService()
}

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServiceComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.StaticFileComponent{},
		&sessions.SessionComponent{},
		handling.NewRouter(
			handling.HandlerEntry{Prefix: "", Handler: store.ProductHandler{}},
			handling.HandlerEntry{Prefix: "", Handler: store.CategoryHandler{}},
			handling.HandlerEntry{Prefix: "", Handler: store.CartHandler{}},
			handling.HandlerEntry{Prefix: "", Handler: store.OrderHandler{}},
			handling.HandlerEntry{Prefix: "", Handler: admin.AdminHandler{}},
			handling.HandlerEntry{Prefix: "", Handler: admin.ProductsHandler{}},
			handling.HandlerEntry{Prefix: "", Handler: admin.CategoriesHandler{}},
			handling.HandlerEntry{Prefix: "", Handler: admin.OrdersHandler{}},
			handling.HandlerEntry{Prefix: "", Handler: admin.DatabaseHandler{}},
		).AddMethodsAlias("/", store.ProductHandler.GetProducts, 0, 1).
			AddMethodsAlias("/products[/]?[A-z0-9]*?", store.ProductHandler.GetProducts, 0, 1).
			AddMethodsAlias("/admin[/]?", admin.AdminHandler.GetSection, ""),
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
