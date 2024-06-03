package admin

import (
	"platform/http/actionresults"
	"platform/http/handling"
	"platform/sessions"
	"sportsstore/models"
)

type ProductsHandler struct {
	models.Repository
	handling.URLGenerator
	sessions.Session
}

type ProductTemplateContext struct {
	Products []models.Product
	EditId   int
	EditUrl  string
	SaveUrl  string
}

const PRODUCT_EDIT_KEY string = "product_edit"

func (handler ProductsHandler) GetData() actionresults.ActionResult {
	return actionresults.NewTemplateAction("admin_products.html", ProductTemplateContext{
		Products: handler.GetProducts(),
		EditId:   handler.Session.GetValueDefault(PRODUCT_EDIT_KEY, 0).(int),
		EditUrl:  mustGenerateURL(handler.URLGenerator, ProductsHandler.PostProductEdit),
		SaveUrl:  mustGenerateURL(handler.URLGenerator, ProductsHandler.PostProductSave),
	})
}

type EditReference struct {
	ID int
}

func (handler ProductsHandler) PostProductEdit(ref EditReference) actionresults.ActionResult {
	handler.Session.SetValue(PRODUCT_EDIT_KEY, ref.ID)
	return actionresults.NewRedirectAction(mustGenerateURL(handler.URLGenerator, AdminHandler.GetSection, "Products"))
}

type ProductSaveReference struct {
	Id                int
	Name, Description string
	Category          int
	Price             float64
}

func (handler ProductsHandler) PostProductSave(
	p ProductSaveReference) actionresults.ActionResult {
	handler.Repository.SaveProduct(&models.Product{
		ID: p.Id, Name: p.Name, Description: p.Description,
		Category: &models.Category{ID: p.Category},
		Price:    p.Price,
	})
	handler.Session.SetValue(PRODUCT_EDIT_KEY, 0)
	return actionresults.NewRedirectAction(mustGenerateURL(handler.URLGenerator,
		AdminHandler.GetSection, "Products"))
}

func mustGenerateURL(gen handling.URLGenerator, target interface{}, data ...interface{}) string {
	url, err := gen.GenerateURL(target, data...)
	if err != nil {
		panic(err)
	}
	return url
}
