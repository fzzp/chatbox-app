package socket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// websocketUpgrader 将HTTP请求升级为持久的websocket连接
var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
		// origin := r.Header.Get("Origin")
		// log.Println("->>> origin: ", origin)
		// switch origin {
		// case "http://localhost:5173": // TODO: 配置
		// 	return true
		// default:
		// 	return false
		// }
	},
}

// SocketServer 管理websocket结构体
type SocketServer struct {
	clients map[*Client]struct{} // 所有链接
	sync.RWMutex
}

// WsEndpoint 初次建立链接握手的才执行此函数，也就是前端执行 new WebSocket("wss://localhost:4321/ws" 的时候。
func (ss *SocketServer) WsEndpoint(c *gin.Context) {
	conn, err := websocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("ip: %s upgrade error: %v\n", c.RemoteIP(), err)
		return
	}

	// 使用 conn 创建一个client
	client := NewClient(conn, ss)
	if ss.clients == nil {
		ss.clients = make(map[*Client]struct{})
	}

	// 添加到 clients 里
	ss.addClient(client)

	// 监听 conn，读写分离
	go client.readJson()
	go client.writeJson()
}

// addClient 添加一个链接
func (ss *SocketServer) addClient(client *Client) {
	ss.Lock()
	defer ss.Unlock()
	ss.clients[client] = struct{}{}
}

// removeClient 移除一个链接
func (ss *SocketServer) removeClient(client *Client) {
	ss.Lock()
	defer ss.Unlock()
	_, ok := ss.clients[client]
	if ok {
		client.conn.Close()
		close(client.payloadChan) // 关闭chan
		delete(ss.clients, client)
	}
}
