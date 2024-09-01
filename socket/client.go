package socket

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ss      *SocketServer
	conn    *websocket.Conn
	msgChan chan Event
}

func NewClient(conn *websocket.Conn, ss *SocketServer) *Client {
	return &Client{
		conn:    conn,
		msgChan: make(chan Event),
		ss:      ss,
	}
}

func (c *Client) readJson() {
	defer func() {
		if r := recover(); r != nil {
			c.ss.removeClient(c)
		}
	}()

	var event Event
	for {
		if err := c.conn.ReadJSON(&event); err != nil {
			log.Printf("%s c.conn.ReadJSON error: %v\n", c.conn.RemoteAddr(), err)
			break
		}
		log.Println("接受到消息：", event.Payload)
		if err := c.ss.dispatchEvent(event, c); err != nil {
			log.Println("dispatchEvent error: ", err)
		}
		// handle, ok := c.ss.handlers[event.Type]
		// if ok {
		// 	if err := handle(event, c); err != nil {
		// 		c.msgChan <- event
		// 	}
		// } else {

		// }
	}
}

func (c *Client) writeJson() {
	defer func() {
		if r := recover(); r != nil {
			c.ss.removeClient(c)
		}
	}()

	for {
		select {
		case event, ok := <-c.msgChan:
			if !ok {
				// TODO: 做点什么
				return
			}

			echo := Event{
				Type:    ET_SendMessage,
				Payload: event.Payload,
			}
			if err := c.conn.WriteJSON(&echo); err != nil {
				log.Printf("c.conn.WriteJSON error: %v, payload: %v", err, echo)
			}
		}
	}
}
