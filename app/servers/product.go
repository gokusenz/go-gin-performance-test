package servers

import (
	"fastwork/go-gin-performance-test/app/helpers"
	"fastwork/go-gin-performance-test/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	contextHelper = &helpers.ContextHelper{}
)

// ProductServer struct
type ProductServer struct {
	ProductService services.IProductService
}

// NewProductServer function to create new product server
func NewProductServer(productService services.IProductService) *ProductServer {
	return &ProductServer{ProductService: productService}
}

// GetByID is product function
func (r *ProductServer) GetByID(c *gin.Context) {

	productID := c.Param("product_id")
	if productID != "" {
		result, err := r.ProductService.GetbyID(productID)

		c.JSON(200, gin.H{
			"message": result,
		})

		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}

}
