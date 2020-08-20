package structs

type UserInfo struct {
	ID int64
	Name string
	Phone string
}

func (info UserInfo) TableName() string {
	return "user_info"
}
