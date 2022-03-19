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

	err := r.db.Select("categories.name as `Category`, brands.name as `Brand`, concat('Rp', format(products.price, 2, 'id_ID')) as `rp`, concat('Rp', format(products.discount, 2, 'id_ID')) as `disc`, products.*").Joins("left join categories on categories.category_id = products.category_id left join brands on brands.brand_id = products.brand_id").Limit(limit).Offset(page).Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) FindByProductID(limit int, page int, productID int) ([]Product, error) {
	var products []Product

	//err := r.db.Debug().Where("product_id = ?", productID).Find(&products)
	err := r.db.Select("categories.name as `Category`, brands.name as `Brand`, concat('Rp', format(products.price, 2, 'id_ID')) as `rp`, concat('Rp', format(products.discount, 2, 'id_ID')) as `disc`, products.*").Joins("left join categories on categories.category_id = products.category_id left join brands on brands.brand_id = products.brand_id").Where("product_id = ?", productID).Limit(limit).Offset(page).Find(&products).Error

	if err != nil {
		return products, err
	}

	return products, nil
}
