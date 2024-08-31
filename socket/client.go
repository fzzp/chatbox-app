package socket

import (
	"chatbox-app/utils"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn        *websocket.Conn
	payloadChan chan WsPayload
	ss          *SocketServer
}

func NewClient(conn *websocket.Conn, ss *SocketServer) *Client {
	return &Client{
		conn:        conn,
		payloadChan: make(chan WsPayload),
		ss:          ss,
	}
}

func (c *Client) readJson() {
	defer func() {
		if r := recover(); r != nil {
			c.ss.removeClient(c)
		}
	}()

	var payload WsPayload
	for {
		if err := c.conn.ReadJSON(&payload); err != nil {
			log.Printf("%s c.conn.ReadJSON error: %v\n", c.conn.RemoteAddr(), err)
			break
		}
		log.Println("接受到消息：", payload.Data)
		// TODO: 抽取封装，处理事件
		c.payloadChan <- payload
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
		case payload, ok := <-c.payloadChan:
			if !ok {
				// TODO: 做点什么
				return
			}

			echo := WsPayload{
				Action: payload.Action,
				Data:   fmt.Sprintf("%v ->>> %d", payload.Data, utils.RandomInt(1, 1000)),
			}
			if err := c.conn.WriteJSON(&echo); err != nil {
				log.Printf("c.conn.WriteJSON error: %v, payload: %v", err, echo)
			}
		}
	}
}
