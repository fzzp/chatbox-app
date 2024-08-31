package models

import "time"

// 验证电子邮件信息表
type VerifyEmail struct {
	BaseField
	UserID     uint      `gorm:"type:int(11);index:idx_userId;not null"`
	Status     uint8     `gorm:"type:int(1);default: 0;not null;comment:'0:正常,1:被使用,2:已失效'"`
	Email      string    `gorm:"type:varchar(96);index:idx_email;not null" json:"email"`
	VerifyCode string    `gorm:"type:varchar(8);index:idx_verify_code;not null" json:"verifyCode"`
	ExpiredAt  time.Time `gorm:"not null" json:"expireAt"`
}
