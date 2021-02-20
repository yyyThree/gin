package entity

import (
	"time"
)

type PrimaryKey struct {
	ID int32
}

type Time struct {
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}
