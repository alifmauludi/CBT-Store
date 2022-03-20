package product

type Service interface {
	GetProducts(limit int, page int, category string, brand string) ([]Product, error)
	GetProductByID(input GetProductDetailInput) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetProducts(limit int, page int, category string, brand string) ([]Product, error) {
	if limit == 0 {
		limit = 100
	}

	if page == 1 {
		page = 0
	} else if page == 2 {
		page = limit
	} else {
		page = limit * (page - 1)
	}

	if category != "" && brand == "" {
		products, err := s.repository.FindByCategory(limit, page, category)
		if err != nil {
			return products, err
		}

		return products, nil
	}

	if category == "" && brand != "" {
		products, err := s.repository.FindByBrand(limit, page, brand)
		if err != nil {
			return products, err
		}

		return products, nil
	}

	if category != "" && brand != "" {
		products, err := s.repository.FindByCatBrand(limit, page, category, brand)
		if err != nil {
			return products, err
		}

		return products, nil
	}

	products, err := s.repository.FindAll(limit, page)
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) GetProductByID(input GetProductDetailInput) (Product, error) {
	product, err := s.repository.FindByID(input.ID)

	if err != nil {
		return product, err
	}

	return product, nil
}
