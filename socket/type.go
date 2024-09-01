package socket

type EventType string

const (
	ET_NewContact  EventType = "new_contact"  // 添加联系人
	ET_SendMessage EventType = "send_message" // 发送信息
)
