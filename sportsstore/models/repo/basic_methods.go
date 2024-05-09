package repo

import "sportsstore/models"

func (repo *SqlRepository) GetProducts() (products []models.Product) {
	rows, err := repo.Commands.GetProducts.QueryContext(repo.Context)
	if err == nil {
		var err error
		if products, err = scanProducts(rows); err != nil {
			repo.Logger.Panicf("cannot scan data: %v", err.Error())
		}
	} else {
		repo.Logger.Panicf("cannot exec GetProduct command: %v", rows.Err().Error())
	}
	return
}

func (repo *SqlRepository) GetProduct(id int) (p models.Product) {
	row := repo.Commands.GetProduct.QueryRowContext(repo.Context, id)
	if row.Err() == nil {
		var err error
		if p, err = scanProduct(row); err != nil {
			repo.Logger.Panicf("cannot scan data: %v", err.Error())
		}
	} else {
		repo.Logger.Panicf("cannot exec GetProduct command: %v", row.Err().Error())
	}
	return
}

func (repo *SqlRepository) GetCategories() []models.Category {
	results := make([]models.Category, 0, 10)
	rows, err := repo.Commands.GetCategories.QueryContext(repo.Context)
	if err == nil {
		for rows.Next() {
			c := models.Category{}
			if err := rows.Scan(&c.ID, &c.CategoryName); err != nil {
				repo.Logger.Panicf("Cannot scan data: %v", err.Error())
			}
			results = append(results, c)
		}
	} else {
		repo.Logger.Panicf("Cannot exec GetCategories command: %v", err)
	}
	return results
}
