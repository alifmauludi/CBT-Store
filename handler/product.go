package handler

import (
	"net/http"
	"store/helper"
	"store/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service product.Service
}

func NewProductHandler(service product.Service) *productHandler {
	return &productHandler{service}
}

func (h *productHandler) GetProducts(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Query("product_id"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	products, err := h.service.GetProducts(limit, page, productID)
	if err != nil {
		response := helper.APIResponse("Error to get products", http.StatusBadRequest, "error", 0, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("List of products", http.StatusOK, "success", len(products), products)
	c.JSON(http.StatusOK, response)
}
