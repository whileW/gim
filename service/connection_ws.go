package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	maxMessageSize = 512
	readDeadline	= 	30*time.Second
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:CheckOrigin,
}
func CheckOrigin(r *http.Request) bool {
	// allow all connections by default
	return true
}
func WebSocketHander(c *gin.Context)  {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("upgrade err:", err)
		return
	}
	conn.SetReadLimit(64*1024)
	conn.SetPongHandler(func(string) error {
		fmt.Println("brower websocket pong...")
		return nil
	})
	fmt.Println("new websocket connection, remote address:", conn.RemoteAddr())

	client := NewClient(new_ws_conn(conn))
	client.Run()
	//handle_client(conn)
}

type ws_conni struct{
	conn 		*websocket.Conn
}

func new_ws_conn(conn *websocket.Conn) *ws_conni {
	ws := &ws_conni{
		conn:conn,
	}
	ws.conn.SetReadLimit(maxMessageSize)
	ws.conn.SetReadDeadline(time.Now().Add(readDeadline))
	return ws
}
func (ws *ws_conni)Read() []byte {
	mt,p,err := ws.conn.ReadMessage()
	if err != nil {
		fmt.Println("read websocket err:", err)
		return nil
	}
	if mt == websocket.TextMessage {
		return p
	}else {
		fmt.Println("invalid websocket message type:", mt)
	}
	return nil
}
func (ws *ws_conni)Write(msg []byte) error {
	w, err := ws.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		fmt.Println("write websocket err:",err)
		return err
	}
	if err := w.Close(); err != nil {
		fmt.Println("write websocket close:"+err.Error())
		return err
	}
	w.Write(msg)
	return nil
}