package models

import "time"

// BaseModel is a base model
type BaseModel struct {
	ID        string `gorm:"Type:uuid;PRIMARY_KEY;DEFAULT:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
