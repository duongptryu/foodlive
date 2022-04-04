package usermodel

import "time"

type Filter struct {
	CreatedAtGt *time.Time
	Phone       string `json:"phone" form:"phone"`
	LastName    string `json:"last_name" form:"last_name"`
}
