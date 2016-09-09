package models

import "time"

type MyModel struct {
	ID        uint `gorm:"primary_key"`
	CreateBy uint
	LastUpdateBy uint
	CreatedAt time.Time
	LastUpdateDate time.Time
	DeletedAt *time.Time `sql:"index"`
}
