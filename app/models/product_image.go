package models

// ProductImage model
type ProductImage struct {
	BaseModel
	ProductID    string `gorm:"INDEX"`
	ImageMedium  string
	IsCoverPhoto bool
}
