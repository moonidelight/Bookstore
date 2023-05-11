package entity

import (
	"time"
)

type Book struct {
	//gorm.Model  `gorm:"softDelete:false"`
	Id          uint   `gorm:"primaryKey"`
	Title       string `gorm:"uniqueIndex"`
	Description string
	Cost        float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	//DeletedAt   gorm.DeletedAt `gorm:"index" sql:"type:index NO ACTION"`
}
