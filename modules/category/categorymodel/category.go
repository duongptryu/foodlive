package categorymodel

import "foodlive/common"

const (
	EntityName = "Category"
)

type Category struct {
	common.SQLModel
	Name        string        `json:"name" gorm:"name"`
	Description string        `json:"description" gorm:"description"`
	Icon        *common.Image `json:"icon" gorm:"icon"`
	Status      bool          `json:"status" gorm:"status"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryCreate struct {
	common.SQLModelCreate
	Name        string        `json:"name" gorm:"name" binding:"required"`
	Description string        `json:"description" gorm:"description"  binding:"required"`
	Icon        *common.Image `json:"icon" gorm:"icon"  binding:"required"`
	Status      bool          `json:"status" gorm:"status"`
}

func (CategoryCreate) TableName() string {
	return Category{}.TableName()
}

func (data *CategoryCreate) Validate() error {
	return nil
}

type CategoryUpdate struct {
	common.SQLModelUpdate
	Name        string        `json:"name" gorm:"name"`
	Description string        `json:"description" gorm:"description"`
	Icon        *common.Image `json:"icon" gorm:"icon"`
	Status      *bool         `json:"status" gorm:"status"`
}

func (CategoryUpdate) TableName() string {
	return Category{}.TableName()
}

func (data *CategoryUpdate) Validate() error {
	if data.Status == nil {
		return ErrStatusCannotNil
	}
	return nil
}

var (
	ErrStatusCannotNil = common.NewCustomError(nil, "Status cannot nil", "ErrStatusCannotNil")
)
