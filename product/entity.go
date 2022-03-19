package product

type Product struct {
	Product_id int    `json:"product_id"`
	Name       string `json:"product_name"`
	Category   string `json:"category"`
	Brand      string `json:"brand"`
	Spec       string `json:"spec"`
	Rp         string `json:"price"`
	Disc       string `json:"discount"`
	Rank       int    `json:"rank"`
}
