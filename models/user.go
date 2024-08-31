package models

// 用户个人信息
type User struct {
	BaseField
	Email        string `gorm:"type:varchar(96);index:idx_email;unique;not null" json:"email"`
	Username     string `gorm:"type:varchar(32);not null" json:"username"`
	PasswordHash string `gorm:"type:varchar(64);not null" json:"-"`
	Avatar       string `gorm:"type:varchar(1024);not null;default:'';" json:"avatar"`
	Intro        string `gorm:"type:varchar(255);not null;default:'这个家伙很懒，什么都没写～';comment:个人简介" json:"intro"`
}
