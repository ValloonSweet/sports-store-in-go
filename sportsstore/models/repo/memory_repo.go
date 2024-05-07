package repo

import (
	"math"
	"platform/services"
	"sportsstore/models"
)

func RegisterMemoryRepoService() {
	services.AddSingleton(func() models.Repository {
		repo := &MemoryRepo{}
		repo.Seed()
		return repo
	})
}

type MemoryRepo struct {
	products   []models.Product
	categories []models.Category
}

// GetProductPageCategory implements models.Repository.
func (r *MemoryRepo) GetProductPageCategory(categoryId int, page int, pageSize int) (products []models.Product, totalAvailable int) {
	if categoryId == 0 {
		return r.GetProductPage(page, pageSize)
	} else {
		filteredProducts := make([]models.Product, 0, len(r.products))
		for _, p := range r.products {
			if p.Category.ID == categoryId {
				filteredProducts = append(filteredProducts, p)
			}
		}
		return getPage(filteredProducts, page, pageSize), len(filteredProducts)
	}
}

// GetProductPage implements models.Repository.
func (r *MemoryRepo) GetProductPage(page int, pageSize int) (products []models.Product, totalAvailable int) {
	return getPage(r.products, page, pageSize), len(r.products)
}

func getPage(src []models.Product, page, pageSize int) []models.Product {
	start := (page - 1) * pageSize
	if page > 0 && len(src) > start {
		end := (int)(math.Min((float64)(len(src)), (float64(start + pageSize))))
		return src[start:end]
	}
	return []models.Product{}
}

// GetCategories implements models.Repository.
func (r *MemoryRepo) GetCategories() []models.Category {
	return r.categories
}

// GetProduct implements models.Repository.
func (r *MemoryRepo) GetProduct(id int) (product models.Product) {
	for _, p := range r.products {
		if p.ID == id {
			product = p
			return
		}
	}
	return
}

// GetProducts implements models.Repository.
func (r *MemoryRepo) GetProducts() []models.Product {
	return r.products
}
