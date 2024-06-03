package admin

import (
	"platform/http/actionresults"
	"sportsstore/models"
)

type CategoriesHandler struct {
	models.Repository
}

func (handler CategoriesHandler) GetData() string {
	return "This is the categories handler"
}

func (handler CategoriesHandler) GetSelect(current int) actionresults.ActionResult {
	return actionresults.NewTemplateAction("select_category.html", struct {
		Current    int
		Categories []models.Category
	}{
		Current: current, Categories: handler.GetCategories(),
	})
}
