package models

// BaseModel is a base model
type BaseModel struct {
	ID string `gorm:"Type:uuid;PRIMARY_KEY;DEFAULT:uuid_generate_v4()"`
}
