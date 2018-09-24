package models

// Product model
type Product struct {
	BaseModel
	Title            string
	Slug             string
	Description      string
	AboutSeller      string
	BasePrice        float64
	Extras           string `gorm:"Type:json;DEFAULT:'{\"options\":[]}'" json:"extras"`
	ExtraDescription string
	ProductImages    ProductImage
}
