package product

import "gorm.io/gorm"

type Repository interface {
	FindAll(limit int, offset int) ([]Product, error)
	FindByProductID(limit, offset, productID int) ([]Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(limit int, page int) ([]Product, error) {
	var products []Product

	err := r.db.Limit(limit).Offset(page).Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) FindByProductID(limit int, page int, productID int) ([]Product, error) {
	var products []Product

	//err := r.db.Debug().Where("product_id = ?", productID).Find(&products)
	err := r.db.Where("product_id = ?", productID).Limit(limit).Offset(page).Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}
