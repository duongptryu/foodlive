package common

type SimpleUser struct {
	SQLModel  `json:",inline"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	Phone     string `json:"phone" gorm:"phone"`
	FirstName string `json:"first_name" gorm:"column:last_name"`
	Role      string `json:"role" gorm:"column:role;"`
	//Avatar    Image  `json:"avatar" gorm:"column:avatar;type:json"`
}

func (SimpleUser) TableName() string {
	return "users"
}
