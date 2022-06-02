package model

type User struct {
	Base
	Name     string `gorm:"type:nvarchar,size:255,not null,default:''"`
	Password string `gorm:"type:nvarchar,size:500,not null,default:''"`
}
