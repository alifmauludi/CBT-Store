package product

type Service interface {
	GetProducts(limit int, page int, productID int) ([]Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetProducts(limit int, page int, productID int) ([]Product, error) {
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

	if productID != 0 {
		products, err := s.repository.FindByProductID(limit, page, productID)
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
