package foodmodel

import "foodlive/common"

const (
	EntityName = "Food"
)

type Food struct {
	common.SQLModel `json:",inline"`
	RestaurantId    int                    `json:"restaurant_id" gorm:"column:restaurant_id"`
	CategoryId      int                    `json:"category_id" gorm:"column:category_id"`
	Category        *common.SimpleCategory `json:"category" gorm:"reference:CategoryId;preload:false"`
	Name            string                 `json:"name" gorm:"column:name"`
	Description     string                 `json:"description" gorm:"column:description"`
	Price           float64                `json:"price" gorm:"column:price"`
	Images          *common.Image          `json:"images" gorm:"images"`
	Status          bool                   `json:"status" gorm:"status"`
	LikeCount       int                    `json:"liked_count" gorm:"column:liked_count"`
	Rating          float64                `json:"rating" gorm:"rating"`
	RatingCount     int                    `json:"rating_count" gorm:"column:rating_count"`
	IsLike          bool                   `json:"is_like" gorm:"-"`
}

func (Food) TableName() string {
	return "foods"
}

type FoodCreate struct {
	common.SQLModelCreate `json:",inline"`
	RestaurantId          int           `json:"restaurant_id" gorm:"column:restaurant_id"`
	CategoryId            int           `json:"category_id" gorm:"column:category_id" binding:"required"`
	Name                  string        `json:"name" gorm:"column:name" binding:"required"`
	Description           string        `json:"description" gorm:"column:description" binding:"required"`
	Price                 float64       `json:"price" gorm:"column:price" binding:"required"`
	Images                *common.Image `json:"images" gorm:"images" binding:"required"`
	Status                bool          `json:"status" gorm:"status"`
}

func (FoodCreate) TableName() string {
	return Food{}.TableName()
}

func (data *FoodCreate) Validate() error {
	if len(data.Name) < 3 {
		return ErrInvalidLengthNameFood
	}

	if len(data.Description) < 20 {
		return ErrInvalidLengthDescriptionFood
	}

	return nil
}

type FoodUpdate struct {
	common.SQLModelUpdate `json:",inline"`
	CategoryId            int           `json:"category_id" gorm:"column:category_id"`
	Name                  string        `json:"name" gorm:"column:name"`
	Description           string        `json:"description" gorm:"column:description"`
	Price                 float64       `json:"price" gorm:"column:price"`
	Images                *common.Image `json:"images" gorm:"images"`
	Status                bool          `json:"status" gorm:"status"`
}

func (FoodUpdate) TableName() string {
	return Food{}.TableName()
}

func (data *FoodUpdate) Validate() error {
	return nil
}
