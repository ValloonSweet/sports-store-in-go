package store

import (
	"platform/http/actionresults"
	"platform/http/handling"
	"sportsstore/models"
)

type CategoryHandler struct {
	Repository   models.Repository
	URLGenerator handling.URLGenerator
}

type categoryTemplateContext struct {
	Categories       []models.Category
	SelectedCategory int
	CategoryUrlFunc  func(int) string
}

func (handler CategoryHandler) GetButtons(selected int) actionresults.ActionResult {
	return actionresults.NewTemplateAction("category_buttons.html",
		categoryTemplateContext{
			Categories:       handler.Repository.GetCategories(),
			SelectedCategory: selected,
			CategoryUrlFunc:  handler.createCategoryFilterFunction(selected),
		})
}

func (handler CategoryHandler) createCategoryFilterFunction(category int) func(int) string {
	return func(category int) string {
		url, _ := handler.URLGenerator.GenerateURL(ProductHandler.GetProducts,
			category, 1)
		return url
	}
}
