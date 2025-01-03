package database

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type InterestsTemplate struct {
	gorm.Model
	ID        uint `gorm:"primaryKey;autoIncrement"`
	AuthorID  uint
	Author    User            `gorm:"foreignKey:AuthorID;references:id"`
	Title     string          `gorm:"type:varchar(100)"`
	Interests json.RawMessage `gorm:"type:json"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
