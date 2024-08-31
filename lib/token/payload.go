package token

import (
	"time"

	"github.com/google/uuid"
)

// Payload 定义Token负载数据
type Payload struct {
	ID        string    `json:"id"`        // Token 唯一标识
	UserID    int64     `json:"userId"`    // 用户id，Token持有者/使用者
	IssuedAt  time.Time `json:"issuedAt"`  // 签发时间
	ExpiredAt time.Time `json:"expiredAt"` // 过期时间
}

// NewPayload 创建一个 Payload 提供给 GenToken 方法使用
func NewPayload(userID int64, duration time.Duration) *Payload {
	payload := &Payload{
		ID:        uuid.NewString(), // 可能会panic
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload
}
