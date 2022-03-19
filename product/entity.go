package product

type Product struct {
	Product_id   int    `json:"product_id"`
	Product_name string `json:"product_name"`
	Category_id  int    `json:"category_id"`
	Brand_id     int    `json:"brand_id"`
	Spec         string `json:"spec"`
	Price        int    `json:"price"`
	Discount     int    `json:"discount"`
	Rank         int    `json:"rank"`
}
