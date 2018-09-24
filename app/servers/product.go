package servers

import (
	"context"
	"fastwork/api/fastwork"
	"fastwork/go-gin-performance-test/app/helpers"
	"fastwork/go-gin-performance-test/app/services"
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
func (r *ProductServer) GetByID(ctx context.Context) (*fastwork.Product, error) {
	// Init variables
	userID := contextHelper.GetData(ctx, "user_id")

	result, err := r.ProductService.GetbyID(in, userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
