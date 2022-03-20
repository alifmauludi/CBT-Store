package product

import "time"

type ProductFormatter struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Category   string    `json:"category"`
	Brand      string    `json:"brand"`
	Rp         string    `json:"price"`
	Disc       string    `json:"discount"`
	Photos     string    `json:"photos"`
	Rank       int       `json:"rank"`
	DateUpdate time.Time `json:"updated_date"`
}

func FormatProduct(product Product) ProductFormatter {
	productFormatter := ProductFormatter{}
	productFormatter.ID = product.ID
	productFormatter.Name = product.Name
	productFormatter.Category = product.Category
	productFormatter.Brand = product.Brand
	productFormatter.Rp = product.Rp
	productFormatter.Disc = product.Disc
	productFormatter.Rank = product.Rank
	productFormatter.DateUpdate = product.DateUpdate
	productFormatter.Photos = ""

	if len(product.Photos) > 0 {
		productFormatter.Photos = "localhost:8080/photos/" + product.Photos[0].FileName
	}

	return productFormatter
}

func FormatProducts(products []Product) []ProductFormatter {
	productsFormatter := []ProductFormatter{}

	for _, product := range products {
		productFormatter := FormatProduct(product)
		productsFormatter = append(productsFormatter, productFormatter)
	}

	return productsFormatter
}

type ProductDetailFormatter struct {
	ID         int                     `json:"id"`
	Name       string                  `json:"name"`
	Category   string                  `json:"category"`
	Brand      string                  `json:"brand"`
	Spec       string                  `json:"spec"`
	Rp         string                  `json:"price"`
	Disc       string                  `json:"discount"`
	Rank       int                     `json:"rank"`
	DateUpdate time.Time               `json:"updated_date"`
	Photos     []ProductPhotoFormatter `json:"photos"`
}

type ProductPhotoFormatter struct {
	FileName string `json:"file_name"`
	IsCover  bool   `json:"is_cover"`
}

func FormatProductDetail(product Product) ProductDetailFormatter {
	productDetailFormatter := ProductDetailFormatter{}
	productDetailFormatter.ID = product.ID
	productDetailFormatter.Name = product.Name
	productDetailFormatter.Category = product.Category
	productDetailFormatter.Brand = product.Brand
	productDetailFormatter.Spec = product.Spec
	productDetailFormatter.Rp = product.Rp
	productDetailFormatter.Disc = product.Disc
	productDetailFormatter.Rank = product.Rank
	productDetailFormatter.DateUpdate = product.DateUpdate

	photos := []ProductPhotoFormatter{}

	for _, photo := range product.Photos {
		photoFormatter := ProductPhotoFormatter{}
		photoFormatter.FileName = "localhost:8080/photos/" + photo.FileName

		isCover := false

		if photo.IsCover == 1 {
			isCover = true
		}
		photoFormatter.IsCover = isCover

		photos = append(photos, photoFormatter)
	}

	productDetailFormatter.Photos = photos

	return productDetailFormatter
}
