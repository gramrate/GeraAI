package database

import (
	"gorm.io/gorm"
	"time"
)

type ConditionTemplate struct {
	gorm.Model
	ID        uint `gorm:"primaryKey;autoIncrement"`
	AuthorID  uint
	Author    User   `gorm:"foreignKey:AuthorID;references:id"`
	Title     string `gorm:"type:varchar(100)"`
	Condition string `gorm:"type:varchar(2000)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
