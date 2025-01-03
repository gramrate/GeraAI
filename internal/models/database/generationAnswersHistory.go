package database

import (
	"gorm.io/gorm"
	"time"
)

type GenerationAnswersHistory struct {
	gorm.Model
	UserID    uint
	User      User   `gorm:"foreignKey:UserID;references:id"`
	Condition string `gorm:"type:varchar(2000)"`
	Answer    string `gorm:"type:varchar(100)"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
