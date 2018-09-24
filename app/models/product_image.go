package models

// ProductImage model
type ProductImage struct {
	BaseModel
	ImageMedium  string
	IsCoverPhoto bool
}
