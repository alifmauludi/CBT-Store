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
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	category := c.Query("category")
	brand := c.Query("brand")

	products, err := h.service.GetProducts(limit, page, category, brand)
	if err != nil {
		response := helper.APIResponse("Error to get products", http.StatusBadRequest, "error", 0, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("List of products", http.StatusOK, "success", len(products), product.FormatProducts(products))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetProduct(c *gin.Context) {
	var input product.GetProductDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of product", http.StatusBadRequest, "error", 0, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productDetail, err := h.service.GetProductByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of product", http.StatusBadRequest, "error", 0, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Product detail", http.StatusOK, "success", 1, product.FormatProductDetail(productDetail))
	c.JSON(http.StatusOK, response)
}
