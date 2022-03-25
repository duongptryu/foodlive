package common

type SimpleRst struct {
	SQLModel `json:",inline"`
	Name     string `json:"name" gorm:"name"`
	Address  string `json:"address" gorm:"column:addr"`
	Logo     *Image `json:"logo" gorm:"logo"`
}

func (SimpleRst) TableName() string {
	return "restaurants"
}
