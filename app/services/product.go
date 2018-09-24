package services

import (
	"fastwork/go-gin-performance-test/app/models"
	"fastwork/go-gin-performance-test/app/repos"

	"github.com/fastworkco/api/fastwork"
)

// ProductService model
type ProductService struct {
	ProductRepo repos.IProductRepo
}

// IProductService interface
type IProductService interface {
	GetbyID(*fastwork.GetProductByIDRequest, string) (*fastwork.Product, error)
}

// NewProductService function to create new product service
func NewProductService(
	productRepo repos.IProductRepo) *ProductService {
	return &ProductService{
		ProductRepo: productRepo,
	}
}

// GetbyID function of product service
func (ps *ProductService) GetbyID(in *fastwork.GetProductByIDRequest, userID string) (*fastwork.Product, error) {
	productRepo := ps.ProductRepo

	// Initial
	referenceID := in.GetReferenceId()

	// Get product
	product, err := productRepo.GetbyID(in.ProductId)
	if err != nil {
		return nil, err
	}

	// If product has packages, use the frist one as base price
	if len(product.ProductPackages) > 0 {
		basePackage := product.ProductPackages[0]
		product.BasePrice = basePackage.Price
	}

	// Convert product to protobuf model
	pbProduct := product.ConvertToProto()

	// Get product order count
	completedOrderCount, err := ps.OrderRepo.CountCompletedByProductID(product.ID)
	if err != nil {
		return nil, err
	}

	// Get product order qeueued
	queuedOrderCount, err := ps.OrderRepo.CountQueuedByProductID(product.ID)
	if err != nil {
		return nil, err
	}

	// Calculate is_new status
	isNew := product.CalculateIsNew(*completedOrderCount, *queuedOrderCount)

	// Get avg rating
	rating, err := ps.ProductReviewRepo.GetAVGRatingByID(product.ID)
	if err != nil {
		return nil, err
	}

	// Prepare Algolia product user model
	pbProduct.User = &fastwork.AlgoliaProductUser{
		Id:                    product.User.ID,
		FirstName:             product.User.FirstName,
		Image:                 product.User.Image,
		LastOnlineAt:          product.User.LastOnlineAt.String(),
		LastOnlineAtTimestamp: int32(product.User.LastOnlineAt.Unix()),
		Username:              product.User.Username,
		Stats: &fastwork.AlgoliaProductUserStat{
			Id:             product.User.UserStat.ID,
			ConversionRate: product.User.UserStat.ConversionRate,
			CompletionRate: product.User.UserStat.CompletionRate,
			ResponseRate:   product.User.UserStat.ResponseRate,
			ResponseTime:   product.User.UserStat.ResponseTime,
		},
		Experience: product.User.Experience,
	}

	// Assign value to proto model
	pbProduct.Rating = rating
	pbProduct.Purchase = *completedOrderCount
	pbProduct.IsNew = isNew

	// Save product action log if user id is exist
	var logUserID *string
	var logReferenceID *string
	if userID != "" {
		logUserID = &userID
	}
	if referenceID != "" {
		logReferenceID = &referenceID
	}

	if logUserID != nil || logReferenceID != nil {
		productActionLog := &models.ProductActionLog{
			UserID:      logUserID,
			ReferenceID: logReferenceID,
			ProductID:   product.ID,
			Action:      "view",
		}
		if err := ps.ProductActionLogRepo.Create(productActionLog); err != nil {
			return nil, err
		}
	}

	return pbProduct, nil
}
