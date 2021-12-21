package common

import "time"

type SQLModel struct {
	Id        int        `json:"id" gorm:"column:id"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at"`
}

type SQLModelCreate struct {
	Id        int        `json:"-" gorm:"column:id"`
	CreatedAt *time.Time `json:"-" gorm:"created_at"`
	UpdatedAt *time.Time `json:"-" gorm:"updated_at"`
}

type SQLModelUpdate struct {
	UpdatedAt *time.Time `json:"-" gorm:"updated_at"`
}
