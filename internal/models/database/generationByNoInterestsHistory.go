package database

import (
	"gorm.io/gorm"
	"time"
)

type GenerationByNoInterestsHistory struct {
	gorm.Model
	UserID    uint
	User      User   `gorm:"foreignKey:UserID;references:id"`
	Condition string `gorm:"type:varchar(2000)"`
	TaskText  string `gorm:"type:varchar(3000)"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
