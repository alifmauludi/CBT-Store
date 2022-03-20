package product

import "gorm.io/gorm"

type Repository interface {
	FindAll(limit int, offset int) ([]Product, error)
	FindByBrand(limit int, offset int, brand string) ([]Product, error)
	FindByCategory(limit int, offset int, category string) ([]Product, error)
	FindByCatBrand(limit int, offset int, category string, brand string) ([]Product, error)
	FindByID(ID int) (Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(limit int, page int) ([]Product, error) {
	var products []Product

	err := r.db.Select("categories.name as `Category`, brands.name as `Brand`, concat('Rp', format(products.price, 2, 'id_ID')) as `rp`, concat('Rp', format(products.discount, 2, 'id_ID')) as `disc`, products.*").Joins("left join categories on categories.id = products.category_id left join brands on brands.id = products.brand_id").Preload("Photos", "photos.is_cover = 1").Limit(limit).Offset(page).Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) FindByBrand(limit int, page int, brand string) ([]Product, error) {
	var products []Product

	//err := r.db.Debug().Where("id = ?", productID).Find(&products)
	err := r.db.Select("categories.name as `Category`, brands.name as `Brand`, concat('Rp', format(products.price, 2, 'id_ID')) as `rp`, concat('Rp', format(products.discount, 2, 'id_ID')) as `disc`, products.*").Joins("left join categories on categories.id = products.category_id left join brands on brands.id = products.brand_id").Where("brands.slug = ?", brand).Preload("Photos", "photos.is_cover = 1").Limit(limit).Offset(page).Find(&products).Error

	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) FindByCategory(limit int, page int, category string) ([]Product, error) {
	var products []Product

	//err := r.db.Debug().Where("id = ?", productID).Find(&products)
	err := r.db.Select("categories.name as `Category`, brands.name as `Brand`, concat('Rp', format(products.price, 2, 'id_ID')) as `rp`, concat('Rp', format(products.discount, 2, 'id_ID')) as `disc`, products.*").Joins("left join categories on categories.id = products.category_id left join brands on brands.id = products.brand_id").Where("categories.slug = ?", category).Preload("Photos", "photos.is_cover = 1").Limit(limit).Offset(page).Find(&products).Error

	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) FindByCatBrand(limit int, page int, category string, brand string) ([]Product, error) {
	var products []Product

	//err := r.db.Debug().Where("id = ?", productID).Find(&products)
	err := r.db.Select("categories.name as `Category`, brands.name as `Brand`, concat('Rp', format(products.price, 2, 'id_ID')) as `rp`, concat('Rp', format(products.discount, 2, 'id_ID')) as `disc`, products.*").Joins("left join categories on categories.id = products.category_id left join brands on brands.id = products.brand_id").Where("categories.slug = ?", category).Where("categories.slug = ? AND brands.slug = ?", category, brand).Preload("Photos", "photos.is_cover = 1").Limit(limit).Offset(page).Find(&products).Error

	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) FindByID(ID int) (Product, error) {
	var product Product
	//err := r.db.Debug().Where("id = ?", productID).Find(&products)
	err := r.db.Select("categories.name as `Category`, brands.name as `Brand`, concat('Rp', format(products.price, 2, 'id_ID')) as `rp`, concat('Rp', format(products.discount, 2, 'id_ID')) as `disc`, products.*").Joins("left join categories on categories.id = products.category_id left join brands on brands.id = products.brand_id").Where("products.id = ?", ID).Preload("Photos").Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}
