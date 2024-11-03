package product

import (
	"net/http"
	"strconv"
	"tech-testing/src/models"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service *ProductService
}

func NewProductController(service *ProductService) *ProductController {
	return &ProductController{service: service}
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := c.service.CreateProduct(ctx, &product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func (c *ProductController) GetProducts(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	skip, _ := strconv.Atoi(ctx.Query("skip"))

	products, err := c.service.GetAllProducts(ctx, limit, skip)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := c.service.UpdateProduct(ctx, uint(id), &product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) RemoveProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.service.RemoveProduct(ctx, uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (c *ProductController) GetProductById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	product, err := c.service.FindProductById(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	ctx.JSON(http.StatusOK, product)
}
