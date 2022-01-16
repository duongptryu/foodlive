package foodmodel

import "foodlive/common"

var (
	ErrInvalidLengthNameFood        = common.NewCustomError(nil, "Name food must be greater than 3 characters", "ErrInvalidLengthNameFood")
	ErrInvalidLengthDescriptionFood = common.NewCustomError(nil, "Description food must be greater than 20 characters", "ErrInvalidLengthDescriptionFood")
	ErrFoodNotFound                 = common.NewCustomError(nil, "Food not found", "ErrFoodNotFound")
)
