package models

type Role struct {
	Id       int
	Title string
	Description  string
	Status    int
	AddTime  int
	//Manager     Manager `gorm:"foreignkey:ManagerId;association_foreignkey:Id"`
}

//定义结构体操作的数据库表
func (Role) TableName() string {
	return "role"
}