package socket

import (
	"encoding/json"
	"log"
)

type Event struct {
	Type    EventType       `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

// EventHandler 这是处理业务事件的签名
type EventHandler func(event Event, c *Client) error

func NewContactHandler(event Event, c *Client) error {
	panic("TODO:")
}

func SendMessageHandler(event Event, c *Client) error {
	log.Println("type: ", event.Type, string(event.Payload))

	// TODO: 根据事件类型做些什么

	// 回个消息
	c.msgChan <- event

	return nil
}
