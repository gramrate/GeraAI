package database

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type GenerationByInterestsHistory struct {
	gorm.Model
	UserID    uint
	User      User            `gorm:"foreignKey:UserID;references:id"`
	Condition string          `gorm:"type:varchar(2000)"`
	Interests json.RawMessage `gorm:"type:json"`
	TaskText  string          `gorm:"type:varchar(3000)"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
