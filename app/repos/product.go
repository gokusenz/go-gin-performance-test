package repos

import (
	"fastwork/go-gin-performance-test/app/models"

	"github.com/jinzhu/gorm"
)

// ProductRepo model
type ProductRepo struct {
	Db *gorm.DB
}

// IProductRepo interface
type IProductRepo interface {
	GetbyID(string) (*models.Product, error)
}

// NewProductRepo function to create new product repo
func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{Db: db}
}

// GetbyID is a dao function to get product by id
func (d ProductRepo) GetbyID(id string) (*models.Product, error) {
	db := d.Db
	var product models.Product
	db = db.Preload("ProductImages", func(db *gorm.DB) *gorm.DB {
		return db.Order("product_images.sort_order ASC")
	})
	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
