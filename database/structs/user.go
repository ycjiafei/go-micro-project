package structs

type UserInfo struct {
	ID int64	`gorm:"primary_key`
	Name string	`gorm:"type:varchar(50)"`
	Phone string `gorm:"type:char(11)"`
}

func (UserInfo) TableName() string {
	return "user_info"
}
