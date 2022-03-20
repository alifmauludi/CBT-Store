package product

type GetProductDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
